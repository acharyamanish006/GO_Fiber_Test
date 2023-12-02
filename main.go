package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/qinains/fastergoding"
)

func main() {
	fastergoding.Run() // hot reload
	app := fiber.New(fiber.Config{
		AppName: "test V1",
	})
	app.Static("/", "./dist")

	app.Get("/get/:id", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("id"))
	})

	log.Fatal(app.Listen(":3000"))

}
