package app

import (
	"log"
	"os"

	"github.com/hedagaurav/blog-api/config"
	"github.com/hedagaurav/blog-api/config/seeder"
	"github.com/hedagaurav/blog-api/routes"
	"github.com/joho/godotenv"
)

var router = routes.SetupRoutes()

func init() {
	// Load the environment variables
	LoadEnv(os.Getenv("ENVIRONMENT"))

	// Initialize the database connection
	config.ConnectDatabase()

	// run seeder
	seeder.SeedData()
}

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

// Run starts the Gin server
func Run() error {
	return router.Run(os.Getenv("APP_URL"))
}
