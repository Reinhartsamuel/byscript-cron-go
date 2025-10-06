# Go Fiber Firebase Redis App

A simple Go Fiber web application with Firebase Firestore and Redis integration.

## Features

- **Fiber Framework**: Fast and lightweight web framework
- **Firebase Firestore**: Cloud database for storing and syncing data
- **Redis**: In-memory data structure store for caching
- **RESTful API**: Simple endpoints for testing integrations

## Prerequisites

- Go 1.21 or later
- Firebase project with Firestore enabled
- Redis server (optional - app will run without it)

## Installation

1. Clone this repository:
```bash
git clone <repository-url>
cd byscript-cron-go
```

2. Install dependencies:
```bash
go mod tidy
```

3. Set up Firebase credentials:
   - Download your Firebase service account key JSON file
   - Set the `FIREBASE_SERVICE_ACCOUNT_JSON` environment variable with the JSON content:
     ```bash
     export FIREBASE_SERVICE_ACCOUNT_JSON='{"type": "service_account", ...}'
     ```
   - Or use the setup script: `./setup-env.sh` which will guide you through the process
   - Use `service-account-key.example.json` as a template for the required format

4. (Optional) Set up Redis:
   - Install and run Redis server locally on port 6379
   - Or use Docker Compose: `docker-compose up -d redis`
   - Or update the Redis connection settings in `main.go`

## Quick Setup

For easy setup, use the provided script:
```bash
./setup-env.sh
```

This will create a `.env` file from the template and provide instructions for setting up Firebase and Redis.

## Configuration

Update the following in `main.go`:

- **Firebase**: Set the environment variable with your service account JSON:
  ```bash
  export FIREBASE_SERVICE_ACCOUNT_JSON='{"type": "service_account", ...}'
  ```

- **Redis**: Update connection settings if needed:
  ```go
  redisClient = redis.NewClient(&redis.Options{
      Addr:     "your-redis-host:6379",
      Password: "your-password",
      DB:       0,
  })
  ```

## Running the Application

### Option 1: Using the run script (recommended)
```bash
./run.sh
```

### Option 2: Manual build and run
```bash
go run main.go
```

### Option 3: Using Docker Compose (includes Redis)
```bash
docker-compose up -d redis
go run main.go
```

The server will start on `http://localhost:3000`

## API Endpoints

### GET /
Basic hello world endpoint
```bash
curl http://localhost:3000/
```

### GET /firebase
Read/write operations with Firebase Firestore
```bash
curl http://localhost:3000/firebase
```

### GET /redis
Read/write operations with Redis
```bash
curl http://localhost:3000/redis
```

## Project Structure

```
byscript-cron-go/
├── main.go                     # Main application file
├── go.mod                      # Go module file
├── go.sum                      # Go dependencies checksum
├── run.sh                      # Run script (recommended)
├── setup-env.sh                # Environment setup script
├── docker-compose.yml          # Redis Docker configuration
├── .env.example                # Environment variables template
├── service-account-key.example.json  # Firebase config template
└── README.md                   # This file
```

## Dependencies

- **github.com/gofiber/fiber/v2**: Web framework
- **firebase.google.com/go**: Firebase SDK
- **cloud.google.com/go/firestore**: Firestore client
- **github.com/go-redis/redis/v8**: Redis client

## Troubleshooting

### Firebase Issues
- Ensure `FIREBASE_SERVICE_ACCOUNT_JSON` environment variable is set with valid JSON
- Check that the service account JSON has proper permissions
- Verify that Firestore is enabled in your Firebase project
- Ensure the Firestore database location matches your configuration

### Redis Issues
- The app will continue running even if Redis is unavailable
- Check that Redis server is running on the specified host and port
- Use `docker-compose up -d redis` to start Redis with Docker
- Verify Redis configuration if using authentication

## License

MIT License