package v1

import (
	"github.com/Axenos-dev/sse-openai-server/internal/controllers/v1/chat"
	"github.com/Axenos-dev/sse-openai-server/internal/controllers/v1/sse"
	"github.com/Axenos-dev/sse-openai-server/internal/llm"
	"github.com/Axenos-dev/sse-openai-server/internal/services"
	"github.com/gofiber/fiber/v2"
)

func InitV1(app *fiber.App) {
	v1 := app.Group("/v1")
	v1.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	llm := llm.LLM{}
	llm.InitClient()

	chat.RegHandlers(v1, services.ChatService{LLM: llm})
	sse.RegHandlers(v1)
}
