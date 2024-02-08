package v1

import (
	"github.com/Axenos-dev/sse-openai-server/db"
	"github.com/Axenos-dev/sse-openai-server/internal/controllers/v1/chat"
	"github.com/Axenos-dev/sse-openai-server/internal/controllers/v1/sse"
	"github.com/Axenos-dev/sse-openai-server/internal/llm"
	"github.com/Axenos-dev/sse-openai-server/internal/services"
	"github.com/Axenos-dev/sse-openai-server/internal/storage"
	"github.com/gofiber/fiber/v2"
)

func InitV1(app *fiber.App, db db.DB) {
	v1 := app.Group("/v1")
	v1.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	llm := llm.LLM{}
	llm.InitClient(storage.NewChatStorage(db))

	chat.RegHandlers(v1, services.ChatService{LLM: llm})
	sse.RegHandlers(v1)
}
