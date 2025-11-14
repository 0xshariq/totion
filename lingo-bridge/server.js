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
    const { text, sourceLocale, targetLocale, fast } = req.body;

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

    // Translate using Lingo.dev SDK
    const result = await lingoDotDev.localizeText(text, {
      sourceLocale: source,
      targetLocale,
      fast: fast || false,
    });

    // Store in cache
    await setCache(cacheKey, result);

    res.json({ translation: result, cached: false });
  } catch (error) {
    console.error('Translation error:', error);
    res.status(500).json({ 
      error: error.message || 'Translation failed' 
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
  console.log(`  GET  /cache/stats        - Cache statistics`);
  console.log(`  POST /cache/clear        - Clear cache`);
});
