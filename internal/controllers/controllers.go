package controllers

import (
	v1 "github.com/Axenos-dev/sse-openai-server/internal/controllers/v1"
	"github.com/gofiber/fiber/v2"
)

func InitControllers(app *fiber.App) {
	v1.InitV1(app)
}
