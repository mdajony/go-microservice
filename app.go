package main

import "github.com/gofiber/fiber/v2"

func main() {
  app := fiber.New()

  app.Get("/hello", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
  })

  app.Get("/greet/:name?", func(c *fiber.Ctx) error {
    if c.Params("name") != "" {
      return c.SendString("Hello " + c.Params("name"))
    }
    return c.SendString("Where is john?")
  })

  app.Listen(":3000")
}
