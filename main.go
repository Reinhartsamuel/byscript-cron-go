package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
)

var (
	firestoreClient *firestore.Client
	redisClient     *redis.Client
	ctx             = context.Background()
)

func main() {
	// Initialize Firebase
	initFirebase()

	// Initialize Redis
	initRedis()

	// Create Fiber app
	app := fiber.New()

	// Basic hello world endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	// Firebase endpoint - read/write data
	app.Get("/firebase", handleFirebase)

	// Redis endpoint - read/write data
	app.Get("/redis", handleRedis)

	// Health check endpoint for Railway
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":    "ok",
			"timestamp": time.Now(),
			"service":   "Go Fiber Firebase Redis App",
		})
	})

	// Get port from environment variable or default to 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Start server
	log.Printf("🚀 Server starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}

func initFirebase() {
	// Get Firebase service account JSON from environment variable
	serviceAccountJSON := os.Getenv("FIREBASE_SERVICE_ACCOUNT_JSON")
	if serviceAccountJSON == "" {
		log.Println("Warning: FIREBASE_SERVICE_ACCOUNT_JSON environment variable not set")
		log.Println("Firebase features will not be available")
		log.Println("Please set FIREBASE_SERVICE_ACCOUNT_JSON with your Firebase service account JSON")
		return
	}

	// Parse the service account JSON
	var serviceAccount map[string]interface{}
	if err := json.Unmarshal([]byte(serviceAccountJSON), &serviceAccount); err != nil {
		log.Printf("Warning: Failed to parse Firebase service account JSON: %v", err)
		log.Println("Continuing without Firebase...")
		return
	}

	// Initialize Firebase with the service account JSON
	opt := option.WithCredentialsJSON([]byte(serviceAccountJSON))
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Printf("Warning: Firebase initialization failed: %v", err)
		log.Println("Continuing without Firebase...")
		return
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Printf("Warning: Firestore client creation failed: %v", err)
		log.Println("Continuing without Firebase...")
		return
	}

	firestoreClient = client
	log.Println("Firebase Firestore initialized successfully")
}

func initRedis() {
	// Initialize Redis client
	// Use Railway Redis environment variables if available, otherwise fallback to localhost
	redisAddr := os.Getenv("REDIS_URL")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}
	
	redisPassword := os.Getenv("REDIS_PASSWORD")
	if redisPassword == "" {
		redisPassword = "" // No password set
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       0,                // Use default DB
	})

	// Test connection
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Printf("Warning: Redis connection failed: %v", err)
		log.Println("Continuing without Redis...")
		redisClient = nil
		return
	}

	log.Println("Redis client initialized successfully")
}

func handleFirebase(c *fiber.Ctx) error {
	if firestoreClient == nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "Firebase service unavailable",
			"message": "Please check Firebase configuration and ensure FIREBASE_SERVICE_ACCOUNT_JSON environment variable is set",
		})
	}

	// Write data to Firestore
	docRef := firestoreClient.Collection("test").Doc("example")
	_, err := docRef.Set(ctx, map[string]interface{}{
		"message":   "Hello from Fiber!",
		"timestamp": time.Now(),
		"source":    "fiber-app",
	})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to write to Firestore: " + err.Error(),
		})
	}

	// Read data from Firestore
	doc, err := docRef.Get(ctx)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to read from Firestore: " + err.Error(),
		})
	}

	var data map[string]interface{}
	if err := doc.DataTo(&data); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to parse Firestore data: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Firebase operation completed",
		"data":    data,
	})
}

func handleRedis(c *fiber.Ctx) error {
	if redisClient == nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "Redis service unavailable",
			"message": "Please ensure Redis server is configured with REDIS_URL environment variable",
		})
	}

	// Write data to Redis
	key := "fiber:test:key"
	value := "Hello from Fiber Redis!"
	err := redisClient.Set(ctx, key, value, 10*time.Minute).Err()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to write to Redis: " + err.Error(),
		})
	}

	// Read data from Redis
	readValue, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to read from Redis: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Redis operation completed",
		"data": fiber.Map{
			"key":   key,
			"value": readValue,
		},
	})
}
