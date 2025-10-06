#!/bin/bash

# Go Fiber Firebase Redis App - Environment Setup Script

set -e

echo "🔧 Setting up environment variables for Go Fiber Firebase Redis App"

# Check if .env file already exists
if [ -f ".env" ]; then
    echo "⚠️  Warning: .env file already exists"
    read -p "Do you want to overwrite it? (y/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        echo "❌ Setup cancelled"
        exit 1
    fi
fi

# Create .env file from example
cp .env.example .env

echo ""
echo "📝 Please edit the .env file with your actual values:"
echo "   nano .env"
echo "   or"
echo "   vim .env"
echo "   or"
echo "   code .env"
echo ""

# Instructions for Firebase setup
echo "🔥 Firebase Setup Instructions:"
echo "1. Go to Firebase Console: https://console.firebase.google.com/"
echo "2. Select your project"
echo "3. Go to Project Settings > Service Accounts"
echo "4. Click 'Generate New Private Key'"
echo "5. Download the JSON file"
echo "6. Copy the entire JSON content"
echo "7. Paste it as the value for FIREBASE_SERVICE_ACCOUNT_JSON in .env"
echo ""

# Instructions for Redis setup
echo "🗄️  Redis Setup (Optional):"
echo "- Uncomment and set REDIS_ADDR if Redis is not on localhost:6379"
echo "- Set REDIS_PASSWORD if your Redis server requires authentication"
echo "- Set REDIS_DB if you want to use a specific database"
echo ""

# Instructions for server configuration
echo "🌐 Server Configuration (Optional):"
echo "- Set PORT if you want to run on a different port than 3000"
echo ""

echo "✅ .env file created from template"
echo "📋 Next steps:"
echo "   1. Edit .env with your actual values"
echo "   2. Run: source .env"
echo "   3. Run: ./run.sh"
echo ""
echo "💡 Tip: You can also set environment variables directly:"
echo "   export FIREBASE_SERVICE_ACCOUNT_JSON='your-json-here'"
echo "   ./run.sh"