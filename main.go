package main

//We’d like you to create a RESTful API with two endpoints.
//1. One of them that fetches the data in the provided MongoDB collection and returns the results
//in the requested format.
//2. Second endpoint is to create(POST) and fetch(GET) data from an in-memory database.
//REQUIREMENTS
//● The code should be written in Golang without using framework. (includes mux, router etc)
//● MongoDB data fetch endpoint should just handle HTTP POST requests.
//DELIVERABLES
//● The public repo URL which has the source code of the project, and a set of
//● instructions if there is any project specific configurations needed to run the project.
//● Public endpoint URLs of the deployed API which is available for testing.

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"getir-case/server"
	"getir-case/service"
	"getir-case/storage"
)

func main() {
	//load env variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbUri := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	collectionName := os.Getenv("COLLECTION_NAME")

	// Initialize MongoDB storage
	mongoStorage, err := storage.NewMongoDBStorage(dbUri, dbName, collectionName)
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB storage: %v", err)
	}
	log.Println("MongoDB storage initialized")

	// Initialize in-memory storage
	inMemoryStorage := storage.NewInMemoryStorage()
	log.Println("In-memory storage initialized")

	//inMemoryStorage := models.NewInMemoryStorage()

	// Initialize API service
	svc := service.NewService(mongoStorage, inMemoryStorage)

	// Initialize HTTP server
	srv := server.NewServer(svc)
	log.Println("HTTP server initialized")

	// Start HTTP server
	srv.Start()
}
