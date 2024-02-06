package chat

import (
	"fmt"
	"net/http"

	"github.com/Axenos-dev/sse-openai-server/internal/entity"
	"github.com/gofiber/fiber/v2"
)

type chat struct {
	chatService ChatService
}

func RegHandlers(r fiber.Router, chatService ChatService) {
	res := chat{chatService}

	r.Post("/chat/:topic", res.sendChatMessage)
}

func (chat chat) sendChatMessage(c *fiber.Ctx) error {
	var req entity.SendChatMessageRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(entity.SendMessageResponse{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("failed to read body: %s", err.Error()),
		})
	}

	if err := req.Validate(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(entity.SendMessageResponse{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("invalid body: %s", err.Error()),
		})
	}

	// chat topic is something like topics in message brokers
	// client can define the topic what it like to listen for,
	// by, for example, taking kind of uuid of the chat where the user
	// is currenttly using
	chatTopic := c.Params("topic")
	if len(chatTopic) == 0 {
		return c.Status(http.StatusBadRequest).JSON(entity.SendMessageResponse{
			Code:    http.StatusBadRequest,
			Message: "empty topic",
		})
	}

	if err := chat.chatService.RunChatCompletionStream(req, chatTopic); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(entity.SendMessageResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: fmt.Sprintf("RunChatCompletionStream failed: %s", err.Error()),
		})
	}

	return c.Status(http.StatusOK).JSON(entity.SendMessageResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("message with content '%s' sent", req.Message),
	})
}
