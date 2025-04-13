package main

import (
	"github.com/hedagaurav/blog-api/config"
	"github.com/hedagaurav/blog-api/routes"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// todo: separate the blog and user code to create separate microservices.
//func init() {
//	// Initialize the database connection
//	config.ConnectDatabase()
//	// Initialize the routes
//	routes.SetupRoutes()
//}

// LoadEnv loads the environment variables from the specified file
func LoadEnv(env string) {
	// load the env file according to the environment
	// if env is empty, load the default env file
	if env == "" {
		env = "development"
	}
	// load the env file
	// if the env file is not found, log the error and exit
	// if the env file is found, load the env variables
	if _, err := os.Stat(".env." + env); os.IsNotExist(err) {
		log.Fatalf("Environment file not found: .env.%s", env)
	} else {
		if err := godotenv.Load(".env." + env); err != nil {
			log.Fatalf("Error loading environment file: %v", err)
		}
	}
}

func main() {
	// Load the environment variables
	LoadEnv(os.Getenv("ENVIRONMENT"))

	// Connect DB
	config.ConnectDatabase()

	// Start the server
	r := routes.SetupRoutes()
	if err := r.Run("localhost:8000"); err != nil {
		panic(err)
	}
}
