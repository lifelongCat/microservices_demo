package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"microservices/internal/api/contact"
	"microservices/internal/api/group"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Microservices Demo",
	})

	api := app.Group("/api/v1")
	api.Route("/contact", contact.CreateRoutes)
	api.Route("/group", group.CreateRoutes)

	if err := app.Listen(":7080"); err != nil {
		log.Fatalln(err)
	}
}
