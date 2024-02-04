package chat

import (
	"fmt"
	"net/http"

	"github.com/Axenos-dev/sse-openai-server/internal/entity"
	"github.com/gin-gonic/gin"
)

type chat struct {
	chatService ChatService
}

func RegHandlers(r *gin.RouterGroup, chatService ChatService) {
	res := chat{chatService}

	r.POST("/chat/:topic", res.sendChatMessage)
}

func (chat chat) sendChatMessage(c *gin.Context) {
	var req entity.SendChatMessageRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, entity.SendMessageResponse{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("failed to read body: %s", err.Error()),
		})
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, entity.SendMessageResponse{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("invalid body: %s", err.Error()),
		})
		return
	}

	// chat topic is something like topics in message brokers
	// client can define the topic what it like to listen for,
	// by, for example, taking kind of uuid of the chat where the user
	// is currenttly using
	chatTopic := c.Param("topic")
	if len(chatTopic) == 0 {
		c.JSON(http.StatusBadRequest, entity.SendMessageResponse{
			Code:    http.StatusBadRequest,
			Message: "empty topic",
		})
		return
	}

	if err := chat.chatService.RunChatCompletionStream(req, chatTopic); err != nil {
		c.JSON(http.StatusUnprocessableEntity, entity.SendMessageResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: fmt.Sprintf("RunChatCompletionStream failed: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, entity.SendMessageResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("message with content '%s' sent", req.Message),
	})
}
