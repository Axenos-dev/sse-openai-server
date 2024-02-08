package services

import (
	"testing"

	"github.com/Axenos-dev/sse-openai-server/internal/entity"
	"github.com/Axenos-dev/sse-openai-server/internal/stream"
	"github.com/stretchr/testify/assert"
)

type mockLLM struct{}

func (mockLLM) StartChatCompletionStream(entity.SendChatMessageRequest, string) {
	// Mock LLM implementation for testing
}

func TestRunChatCompletionStream(t *testing.T) {
	chatService := ChatService{
		LLM: mockLLM{},
	}

	topic := "6"
	req := entity.SendChatMessageRequest{
		Message: "Test message",
	}

	defer stream.MessageCompletionStream.CloseStream(topic)

	// Case 1: Stream exists, expected successful stream initiation
	t.Run("Success", func(t *testing.T) {
		stream.MessageCompletionStream.InitNewStream(topic)
		err := chatService.RunChatCompletionStream(req, topic)

		assert.NoError(t, err)
	})

	// Case 2: Stream does not exist, expected error
	t.Run("Failure_No_Stream", func(t *testing.T) {
		stream.MessageCompletionStream.CloseStream(topic)
		err := chatService.RunChatCompletionStream(req, topic)

		assert.Error(t, err)
	})
}
