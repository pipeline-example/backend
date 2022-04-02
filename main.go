package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Data struct {
  Name string `json:"name"`
  Age int `json:"age"`
}

func main() {
  app := fiber.New()
  app.Use(cors.New())

  app.Get("/", func(c *fiber.Ctx) error {
    data := Data{
      Name: "Chalvin",
      Age: 20,
    }
    return c.JSON(data)
  })

  app.Listen(":80")
}