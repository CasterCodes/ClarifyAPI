package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from Fiber on EC2 🚀")
	})

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello route")
	})

	log.Fatal(app.Listen(":8080"))

}