package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		resp, err := http.Get("http://localhost:8092/surveys/sample")
		if err != nil {
			fmt.Println(err)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		return c.SendString(string(body))
	})

	app.Listen(":8080")
}
