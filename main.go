package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pusher/pusher-http-go/v5"
)

func main() {
    app := fiber.New()

	app.Use(cors.New())

	pusherClient := pusher.Client{
		AppID: "1552215",
		Key: "92305e3b9a099daaa33f",
		Secret: "bf6bf7d168dd04bc8d56",
		Cluster: "ap1",
		Secure: true,
	}

    app.Post("/", func(c *fiber.Ctx) error {

		var data map[string]string

		if err := c.BodyParser(&data); err != nil {
			return err
		}

		pusherClient.Trigger("chat", "massage", data)
        return c.JSON([]string{})
    })

    app.Listen(":8000")
}