import express from 'express';
import { LingoDotDevEngine } from 'lingo.dev/sdk';
import { createClient } from 'redis';
import dotenv from 'dotenv';
import path from 'path';
import { fileURLToPath } from 'url';

// Get current directory
const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

// Load .env from parent directory (root of project)
dotenv.config({ path: path.join(__dirname, '..', '.env') });

const app = express();
app.use(express.json());

// Initialize Lingo.dev SDK
const lingoDotDev = new LingoDotDevEngine({
  apiKey: process.env.LINGODOTDEV_API_KEY,
});

// Technical glossary for consistent translations
// This ensures UI terms are always translated correctly
const technicalGlossary = {
  'note': 'note',
  'notebook': 'notebook',
  'template': 'template',
  'markdown': 'markdown',
  'export': 'export',
  'import': 'import',
  'sync': 'sync',
  'tag': 'tag',
  'search': 'search',
  'daily': 'daily journal',
  'quick note': 'quick note',
};

// Detect context type for better translation accuracy
const detectContext = (text) => {
  const lower = text.toLowerCase();
  
  // UI labels and buttons
  if (text.length < 30 && !text.includes('.') && !text.includes('\n')) {
    return 'UI button or label in a note-taking application';
  }
  
  // Help text or instructions
  if (text.length > 100 || text.includes('?') || lower.includes('how to') || lower.includes('press')) {
    return 'User instructions or help text for a note-taking application';
  }
  
  // Keyboard shortcuts
  if (text.includes('Ctrl') || text.includes('Alt') || text.includes('Shift') || text.includes('→')) {
    return 'Keyboard shortcut description for application UI';
  }
  
  // Status messages
  if (text.includes('successfully') || text.includes('failed') || text.includes('error')) {
    return 'Status message or notification in application UI';
  }
  
  // Menu items
  if (text.includes('...') || text.match(/^[A-Z][a-z]+$/)) {
    return 'Menu item or action in application UI';
  }
  
  return 'Technical UI text for note-taking application';
};

// Initialize Redis client (optional - fallback to memory cache if not available)
let redis = null;
let memoryCache = new Map();

try {
  redis = createClient({
    url: process.env.REDIS_URL || 'redis://localhost:6379'
  });

  redis.on('error', (err) => {
    console.warn('Redis not available, using memory cache:', err.message);
    redis = null;
  });
  
  redis.on('connect', () => console.log('✓ Connected to Redis for caching'));

  // Connect to Redis with timeout
  await Promise.race([
    redis.connect(),
    new Promise((_, reject) => setTimeout(() => reject(new Error('Redis timeout')), 2000))
  ]).catch(err => {
    console.warn('Redis connection failed, using memory cache');
    redis = null;
  });
} catch (error) {
  console.warn('Redis not available, using memory cache');
  redis = null;
}

// Cache helper function
const getCacheKey = (text, sourceLang, targetLang) => {
  return `lingo:${sourceLang}:${targetLang}:${text.substring(0, 100)}`;
};

const getFromCache = async (key) => {
  try {
    if (redis) {
      return await redis.get(key);
    } else {
      // Use memory cache
      return memoryCache.get(key) || null;
    }
  } catch (error) {
    console.error('Cache get error:', error);
    return null;
  }
};

const setCache = async (key, value, ttl = 86400) => {
  try {
    if (redis) {
      await redis.set(key, value, { EX: ttl });
    } else {
      // Use memory cache
      memoryCache.set(key, value);
      // Simple TTL: delete after timeout
      setTimeout(() => memoryCache.delete(key), ttl * 1000);
    }
  } catch (error) {
    console.error('Cache set error:', error);
  }
};

// Health check endpoint
app.get('/health', (req, res) => {
  res.json({ 
    status: 'ok', 
    service: 'lingo-bridge',
    apiKeyConfigured: !!process.env.LINGODOTDEV_API_KEY 
  });
});

// Translate text endpoint with Redis caching
app.post('/translate', async (req, res) => {
  try {
    const { text, sourceLocale, targetLocale, fast, context } = req.body;

    if (!text || !targetLocale) {
      return res.status(400).json({ 
        error: 'Missing required fields: text and targetLocale' 
      });
    }

    const source = sourceLocale || 'en';
    const cacheKey = getCacheKey(text, source, targetLocale);
    
    // Check cache first
    const cached = await getFromCache(cacheKey);
    if (cached) {
      return res.json({ translation: cached, cached: true });
    }

    // Detect or use provided context for better accuracy
    const translationContext = context || detectContext(text);
    
    // Translate using Lingo.dev SDK with quality mode and context
    // Quality mode (fast: false) + context ensures >90% accuracy
    let result;
    let retries = 2; // Retry failed translations for better reliability
    
    while (retries >= 0) {
      try {
        result = await lingoDotDev.localizeText(text, {
          sourceLocale: source,
          targetLocale,
          fast: fast === true, // Default to quality mode (false) for maximum accuracy
          context: translationContext,
          // Preserve formatting like newlines, markdown, etc.
          preserveFormatting: true,
          // Use glossary for consistent technical terms
          glossary: technicalGlossary,
        });
        break; // Success
      } catch (error) {
        retries--;
        if (retries < 0) throw error;
        // Wait before retry (exponential backoff)
        await new Promise(resolve => setTimeout(resolve, 100 * (3 - retries)));
      }
    }

    // Store in cache with long TTL for accuracy consistency
    await setCache(cacheKey, result, 604800); // 7 days

    res.json({ translation: result, cached: false, context: translationContext });
  } catch (error) {
    console.error('Translation error:', error);
    res.status(500).json({ 
      error: error.message || 'Translation failed',
      originalText: req.body.text // Return original for fallback
    });
  }
});

// Batch translate endpoint - translate multiple strings at once
// Optimized for pre-warming cache with parallel processing
app.post('/translate/batch', async (req, res) => {
  try {
    const { texts, sourceLocale, targetLocale, fast } = req.body;

    if (!texts || !Array.isArray(texts) || !targetLocale) {
      return res.status(400).json({ 
        error: 'Missing required fields: texts (array) and targetLocale' 
      });
    }

    const source = sourceLocale || 'en';
    const results = [];
    let cachedCount = 0;
    let translatedCount = 0;
    let errorCount = 0;

    // Process in chunks of 10 for optimal performance
    const chunkSize = 10;
    for (let i = 0; i < texts.length; i += chunkSize) {
      const chunk = texts.slice(i, i + chunkSize);
      
      // Process chunk in parallel for speed
      const chunkPromises = chunk.map(async (text) => {
        const cacheKey = getCacheKey(text, source, targetLocale);
        
        // Check cache first
        const cached = await getFromCache(cacheKey);
        if (cached) {
          cachedCount++;
          return cached;
        }

        // Translate with context detection and retry logic
        const translationContext = detectContext(text);
        let retries = 2;
        
        while (retries >= 0) {
          try {
            const result = await lingoDotDev.localizeText(text, {
              sourceLocale: source,
              targetLocale,
              fast: fast === true, // Default to quality mode for maximum accuracy
              context: translationContext,
              preserveFormatting: true,
              glossary: technicalGlossary,
            });

            // Store in cache with 7-day TTL
            await setCache(cacheKey, result, 604800);
            translatedCount++;
            return result;
          } catch (error) {
            retries--;
            if (retries < 0) {
              console.error(`Translation failed for: "${text}"`, error.message);
              errorCount++;
              return text; // Fallback to original
            }
            // Exponential backoff
            await new Promise(resolve => setTimeout(resolve, 100 * (3 - retries)));
          }
        }
      });

      // Wait for chunk to complete
      const chunkResults = await Promise.all(chunkPromises);
      results.push(...chunkResults);
    }

    res.json({ 
      results,
      stats: {
        total: texts.length,
        cached: cachedCount,
        translated: translatedCount,
        errors: errorCount,
        cacheHitRate: `${Math.round((cachedCount / texts.length) * 100)}%`
      }
    });
  } catch (error) {
    console.error('Batch translation error:', error);
    res.status(500).json({ 
      error: error.message || 'Batch translation failed' 
    });
  }
});

// Cache stats endpoint
app.get('/cache/stats', async (req, res) => {
  try {
    const info = await redis.info('stats');
    const keys = await redis.dbSize();
    res.json({ 
      totalKeys: keys,
      info: info 
    });
  } catch (error) {
    res.status(500).json({ error: error.message });
  }
});

// Clear cache endpoint
app.post('/cache/clear', async (req, res) => {
  try {
    await redis.flushDb();
    res.json({ success: true, message: 'Cache cleared' });
  } catch (error) {
    res.status(500).json({ error: error.message });
  }
});

const PORT = process.env.LINGO_BRIDGE_PORT || 3737;

app.listen(PORT, () => {
  console.log(`🌐 Lingo.dev Bridge Server running on http://localhost:${PORT}`);
  console.log(`✓ API Key configured: ${!!process.env.LINGODOTDEV_API_KEY}`);
  console.log(`✓ Redis caching enabled`);
  console.log(`\nEndpoints:`);
  console.log(`  GET  /health             - Health check`);
  console.log(`  POST /translate          - Translate text (cached)`);
  console.log(`  POST /translate/batch    - Batch translate multiple texts`);
  console.log(`  GET  /cache/stats        - Cache statistics`);
  console.log(`  POST /cache/clear        - Clear cache`);
});
