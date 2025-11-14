# Lingo.dev Bridge Server

This is a Node.js bridge server that provides a REST API interface to the Lingo.dev JavaScript SDK.

## Why?

Lingo.dev only provides a JavaScript SDK, not a REST API. This bridge server allows the Go application to use Lingo.dev's translation services by exposing the SDK through HTTP endpoints.

## Setup

1. Install dependencies:
```bash
cd lingo-bridge
npm install
```

2. Make sure your `.env` file in the parent directory has the Lingo.dev API key:
```bash
LINGODOTDEV_API_KEY=your_api_key_here
```

3. Start the server:
```bash
npm start
```

The server will run on `http://localhost:3737` by default.

## API Endpoints

### Health Check
```bash
GET /health
```

### Translate Text
```bash
POST /translate
Content-Type: application/json

{
  "text": "Hello, world!",
  "sourceLocale": "en",
  "targetLocale": "es",
  "fast": false
}

Response: { "translation": "¡Hola Mundo!" }
```

### Batch Translate
```bash
POST /translate/batch
Content-Type: application/json

{
  "text": "Hello, world!",
  "sourceLocale": "en",
  "targetLocales": ["es", "fr", "de"]
}

Response: { "translations": ["¡Hola Mundo!", "Bonjour le monde!", "Hallo Welt!"] }
```

### Translate Object
```bash
POST /translate/object
Content-Type: application/json

{
  "object": {
    "greeting": "Hello",
    "farewell": "Goodbye"
  },
  "sourceLocale": "en",
  "targetLocale": "es"
}

Response: { "translation": { "greeting": "Hola", "farewell": "Adiós" } }
```

### Detect Language
```bash
POST /detect
Content-Type: application/json

{
  "text": "Bonjour le monde"
}

Response: { "locale": "fr" }
```

## Running in Production

Keep the bridge server running alongside your Go application:

```bash
# Terminal 1: Start bridge server
cd lingo-bridge
npm start

# Terminal 2: Run Go application
cd ..
make run
```

Or use a process manager like PM2:

```bash
npm install -g pm2
pm2 start server.js --name lingo-bridge
pm2 save
```
