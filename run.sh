#!/bin/bash

# Go Fiber Firebase Redis App - Run Script

set -e

echo "🚀 Starting Go Fiber Firebase Redis App..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.21 or later."
    exit 1
fi

# Check if Firebase service account environment variable is set
if [ -z "$FIREBASE_SERVICE_ACCOUNT_JSON" ]; then
    echo "⚠️  Warning: FIREBASE_SERVICE_ACCOUNT_JSON environment variable not set"
    echo "   Firebase features will not be available"
    echo "   Please set FIREBASE_SERVICE_ACCOUNT_JSON with your Firebase service account JSON"
fi

# Check if Redis is running (optional)
if ! command -v redis-cli &> /dev/null; then
    echo "⚠️  Warning: redis-cli not found, cannot check Redis server"
    echo "   Redis features may not be available"
else
    if redis-cli ping &> /dev/null; then
        echo "✅ Redis server is running"
    else
        echo "⚠️  Warning: Redis server is not responding"
        echo "   Redis features will not be available"
        echo "   You can start Redis with: docker-compose up -d redis"
    fi
fi

# Install dependencies if needed
echo "📦 Checking dependencies..."
go mod tidy

# Build the application
echo "🔨 Building application..."
go build -o app

# Run the application
echo "🌐 Starting server on http://localhost:3000"
echo "   Endpoints:"
echo "   - GET /         - Hello World"
echo "   - GET /firebase - Firebase operations"
echo "   - GET /redis    - Redis operations"
echo ""
echo "Press Ctrl+C to stop the server"

./app
