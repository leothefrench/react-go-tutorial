package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println(("Hello, Léo bien sûr"))
	app := fiber.New()

	app.Get("/", func(c *fiber.ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Hello Léo"})
	})

	log.Fatal(app.Listen(":4000"))
}