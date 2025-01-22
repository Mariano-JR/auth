package main

import (
	"log"
	"os"

	"github.com/Mariano-JR/auth/cmd/routes"
	"github.com/Mariano-JR/auth/internal/db"
	"github.com/Mariano-JR/auth/internal/user"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	db.Connect()

	if err := db.DB.AutoMigrate(&user.User{}); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	app := fiber.New()

	routes.Setup(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}
