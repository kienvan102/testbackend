package main

import (
	"log"
	"os"
	"testbackend/handlers"
	"testbackend/repositories"

	"github.com/gin-gonic/gin"
 	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

    // Set default values if environment variables are not found
    dbUser := getEnv("DB_USER", "root")
    dbPassword := getEnv("DB_PASSWORD", "mypassword")
    dbName := getEnv("DB_NAME", "kaban")
    dbHost := getEnv("DB_HOST", "localhost")
    dbPort := getEnv("DB_PORT", "5432")
    db, err := repositories.InitDB(dbUser, dbPassword, dbName, dbHost, dbPort)
    if err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }
    defer db.Close()

    // Initialize handlers with database connection
    handlers.NewTicketHandler(db)

    r := gin.Default()
	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))
    r.GET("/tickets", handlers.GetTickets)
    r.POST("/tickets", handlers.CreateTicket)

    r.Run(":8080") 
}

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}