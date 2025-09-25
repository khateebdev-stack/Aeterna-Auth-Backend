package main

import (
	"aeterna-auth/internal/api/handlers" // یہ نئی import ہے
	"aeterna-auth/internal/database"
	"aeterna-auth/internal/services" // یہ نئی import ہے
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// The main function is the entry point of the program
func main() {
	// Loading environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error while loading .env file")
	}

	// Now, we get the values from the environment variables
	pgConnStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DBNAME"))

	redisAddr := os.Getenv("REDIS_ADDR")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	// Initialize PostgreSQL and Redis connections
	if err := database.InitPostgresDB(pgConnStr); err != nil { //InitPostgresDb is basicaly a function in /internal/database we called it here to pass our DB connection string ,
		log.Fatalf("Could not connect to Database : %v ", err)
	}

	defer database.ClosePostgresDB()

	if err := database.InitRedisClient(redisAddr, redisPassword); err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	// یہاں نیا کوڈ شامل کرنا ہے
	// Initialize services
	userService := services.NewUserService(database.DB)

	// Create a new router
	router := mux.NewRouter()

	// Define API routes.
	router.HandleFunc("/api/register", handlers.RegisterUser(userService)).Methods("POST")
	router.HandleFunc("/api/login", handlers.LoginUser(userService)).Methods(("POST"))
	router.HandleFunc("/api/user/delete", handlers.DeleteUser(userService)).Methods("Delete")
	router.HandleFunc("/api/user/update-password", handlers.UpdatePassword(userService)).Methods("PUT")
	router.HandleFunc("/api/user/update-email", handlers.UpdateEmail(userService)).Methods("PUT")

	// Old route
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to AeternaAuth Authentication system")
	}).Methods("GET")

	// Start the HTTP server
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
