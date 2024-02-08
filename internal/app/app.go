package app

import (
	"fmt"

	"github.com/Axenos-dev/sse-openai-server/config"
	"github.com/Axenos-dev/sse-openai-server/db"
	"github.com/Axenos-dev/sse-openai-server/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

func Run() error {
	db, err := db.NewDB(config.Config.PostgreSQL)
	if err != nil {
		return fmt.Errorf("storage initialization error: %v", err)
	}

	defer db.Close()

	if err := db.Migrate(); err != nil {
		return fmt.Errorf("migration error: %v", err)
	}

	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"message": "pong",
		})
	})

	controllers.InitControllers(app, db)

	return app.Listen(fmt.Sprintf(":%s", config.Config.Port))
}
