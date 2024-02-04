package chat

import "github.com/Axenos-dev/sse-openai-server/internal/entity"

type ChatService interface {
	RunChatCompletionStream(entity.SendChatMessageRequest, string) error
}
