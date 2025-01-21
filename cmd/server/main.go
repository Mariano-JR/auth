package main

import (
	"log"
	"os"

	"github.com/Mariano-JR/auth/cmd/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	app := fiber.New()

	routes.Setup(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}
