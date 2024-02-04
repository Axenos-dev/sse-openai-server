package llm

import (
	"context"
	"errors"
	"io"
	"log"
	"time"

	"github.com/Axenos-dev/sse-openai-server/config"
	"github.com/Axenos-dev/sse-openai-server/internal/entity"
	"github.com/Axenos-dev/sse-openai-server/internal/stream"
	openai "github.com/sashabaranov/go-openai"
)

type LLM struct {
	c *openai.Client
}

func (s *LLM) InitClient() {
	s.c = openai.NewClient(config.Config.OpenAI.ApiKey)
}

func (llm LLM) StartChatCompletionStream(message entity.SendChatMessageRequest, topic string) {
	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: int(config.Config.OpenAI.MaxTokens),
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: message.Message,
			},
		},
		Stream: true,
	}

	s, err := llm.c.CreateChatCompletionStream(context.Background(), req)
	if err != nil {
		log.Printf("ChatCompletionStream error: %v\n", err)
		return
	}
	defer s.Close()

	contentStack := ""
	for {
		select {
		case <-stream.MessageCompletionStream.EndChan(topic):
			log.Printf("Stop signal for '%s'\n", topic)
			stream.MessageCompletionStream.Write(entity.MessageCompletionStream{
				Event: entity.StreamEventEndOfMsg,
				Data: entity.StreamData{
					Content: contentStack,
				},
				Topic: topic,
			}, topic)
			return

		default:
			response, err := s.Recv()
			if errors.Is(err, io.EOF) {
				log.Printf("Stream finished for '%s'\n", topic)
				stream.MessageCompletionStream.Write(entity.MessageCompletionStream{
					Event: entity.StreamEventEndOfMsg,
					Data: entity.StreamData{
						Content: contentStack,
					},
					Topic: topic,
				}, topic)
				return
			}

			if err != nil {
				log.Printf("\nStream error for '%s': %v\n", topic, err)
				stream.MessageCompletionStream.Write(entity.MessageCompletionStream{
					Event: entity.StreamEventError,
					Data: entity.StreamData{
						Content: contentStack,
					},
					Topic:   topic,
					Message: err.Error(),
				}, topic)
				return
			}

			contentStack += response.Choices[0].Delta.Content

			err = stream.MessageCompletionStream.Write(entity.MessageCompletionStream{
				Event: entity.StreamEventMessageComletion,
				Data: entity.StreamData{
					Content: contentStack,
				},
				Topic: topic,
			}, topic)

			if err != nil {
				log.Printf("%v - ending stream for '%s'\n", err, topic)
				return
			}

			//some delay for visibility
			time.Sleep(time.Millisecond * 700)
		}
	}
}
