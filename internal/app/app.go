package app

import (
	"fmt"

	"github.com/Axenos-dev/sse-openai-server/config"
	"github.com/Axenos-dev/sse-openai-server/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

func Run() error {
	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"message": "pong",
		})
	})

	controllers.InitControllers(app)

	return app.Listen(fmt.Sprintf(":%s", config.Config.Port))
}
