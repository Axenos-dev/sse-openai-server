package llm

import (
	"context"

	"github.com/Axenos-dev/sse-openai-server/internal/entity"
)

type chatStorage interface {
	CreateChatMessage(ctx context.Context, message string, msg entity.MessageCompletionStream) error
}
