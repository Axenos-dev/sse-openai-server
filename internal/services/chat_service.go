package services

import (
	"fmt"

	"github.com/Axenos-dev/sse-openai-server/internal/entity"
	"github.com/Axenos-dev/sse-openai-server/internal/stream"
)

type LLM interface {
	StartChatCompletionStream(req entity.SendChatMessageRequest, topic string)
}

type ChatService struct {
	LLM LLM
}

func (s ChatService) RunChatCompletionStream(req entity.SendChatMessageRequest, topic string) error {
	stream.MessageCompletionStream.EndStream(topic)

	// i decided to stream only for one who listening
	// if not listening there, so no clue to stream for
	if stream.MessageCompletionStream.DoesStreamExist(topic) {
		go s.LLM.StartChatCompletionStream(req, topic)
	} else {
		return fmt.Errorf("cant stream to '%s'. That stream do not exists", topic)
	}

	return nil
}
