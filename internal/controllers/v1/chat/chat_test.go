package chat

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Axenos-dev/sse-openai-server/internal/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type mockChatService struct{}

func (mockChatService) RunChatCompletionStream(entity.SendChatMessageRequest, string) error {
	return nil
}

func TestSendChatMessage(t *testing.T) {
	app := fiber.New()
	chat := chat{mockChatService{}}

	app.Post("/v1/chat/:topic", chat.sendChatMessage)

	// Test case 1: Valid request
	t.Run("SendChatMessage_Success", func(t *testing.T) {
		reqBody := entity.SendChatMessageRequest{Message: "Hello!"}
		reqJSON, _ := json.Marshal(reqBody)

		req := httptest.NewRequest("POST", "/v1/chat/1", bytes.NewBuffer(reqJSON))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		defer resp.Body.Close()

		var response entity.SendMessageResponse
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	// Test case 2: Invalid request
	t.Run("SendChatMessage_Failure", func(t *testing.T) {
		reqBody := entity.SendChatMessageRequest{Message: "      "}
		reqJSON, _ := json.Marshal(reqBody)

		req := httptest.NewRequest("POST", "/v1/chat/1", bytes.NewBuffer(reqJSON))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		defer resp.Body.Close()

		var response entity.SendMessageResponse
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
}
