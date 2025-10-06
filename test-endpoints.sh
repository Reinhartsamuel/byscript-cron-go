#!/bin/bash

# Go Fiber Firebase Redis App - Endpoint Test Script

set -e

echo "🧪 Testing Go Fiber Firebase Redis App Endpoints"

# Base URL - change this to your deployed URL or keep localhost for local testing
BASE_URL="${1:-http://localhost:3000}"

echo "Testing endpoints on: $BASE_URL"
echo ""

# Function to test endpoint with pretty output
test_endpoint() {
    local endpoint=$1
    local description=$2
    local url="$BASE_URL$endpoint"
    
    echo "🔍 Testing: $description"
    echo "   URL: $url"
    
    # Make the request and capture response
    response=$(curl -s -w "\n%{http_code}" "$url")
    http_code=$(echo "$response" | tail -n1)
    body=$(echo "$response" | head -n -1)
    
    echo "   Status: $http_code"
    
    if [ "$http_code" -eq 200 ]; then
        echo "   ✅ Success"
        # Pretty print JSON response if available
        if command -v jq &> /dev/null && [ -n "$body" ]; then
            echo "   Response:"
            echo "$body" | jq '.' 2>/dev/null || echo "   $body"
        else
            echo "   Response: $body"
        fi
    else
        echo "   ❌ Failed with HTTP $http_code"
        echo "   Response: $body"
    fi
    echo ""
}

# Test all endpoints
test_endpoint "/" "Hello World endpoint"
test_endpoint "/health" "Health check endpoint"
test_endpoint "/firebase" "Firebase operations endpoint"
test_endpoint "/redis" "Redis operations endpoint"

echo "🎉 All endpoint tests completed!"
echo ""
echo "📝 Summary:"
echo "   If Firebase and Redis are not configured, those endpoints will return 503 status"
echo "   This is expected behavior - the app continues running without these services"
echo ""
echo "💡 Next steps:"
echo "   1. Set FIREBASE_SERVICE_ACCOUNT_JSON environment variable for Firebase"
echo "   2. Add Redis service or set REDIS_URL for Redis functionality"
echo "   3. For Railway deployment, set these as environment variables in your project settings"