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
	c           *openai.Client
	chatStorage chatStorage
}

func (s *LLM) InitClient(chatStorage chatStorage) {
	s.c = openai.NewClient(config.Config.OpenAI.ApiKey)
	s.chatStorage = chatStorage
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

			msg := entity.MessageCompletionStream{
				Event: entity.StreamEventEndOfMsg,
				Data: entity.StreamData{
					Content: contentStack,
				},
				Topic:  topic,
				Status: entity.MessageStatusInterrupted,
			}

			if err := stream.MessageCompletionStream.Write(msg, topic); err != nil {
				log.Printf("write to chan error for '%s': %v", topic, err)
			}

			if contentStack != "" {
				if err := llm.chatStorage.CreateChatMessage(context.Background(), message.Message, msg); err != nil {
					log.Printf("failed to create chat message in db: %v", err)
				}
			}

			return

		default:
			response, err := s.Recv()
			if errors.Is(err, io.EOF) {
				log.Printf("Stream finished for '%s'\n", topic)

				msg := entity.MessageCompletionStream{
					Event: entity.StreamEventEndOfMsg,
					Data: entity.StreamData{
						Content: contentStack,
					},
					Topic:   topic,
					Message: err.Error(),
					Status:  entity.MessageStatusDone,
				}

				if err := stream.MessageCompletionStream.Write(msg, topic); err != nil {
					log.Printf("write to chan error for '%s': %v", topic, err)
				}

				if contentStack != "" {
					if err := llm.chatStorage.CreateChatMessage(context.Background(), message.Message, msg); err != nil {
						log.Printf("failed to create chat message in db: %v", err)
					}
				}

				return
			}

			if err != nil {
				log.Printf("\nStream error for '%s': %v\n", topic, err)

				msg := entity.MessageCompletionStream{
					Event: entity.StreamEventError,
					Data: entity.StreamData{
						Content: contentStack,
					},
					Topic:   topic,
					Message: err.Error(),
					Status:  entity.MessageStatusInterrupted,
				}

				if err := stream.MessageCompletionStream.Write(msg, topic); err != nil {
					log.Printf("write to chan error for '%s': %v", topic, err)
				}

				if contentStack != "" {
					if err := llm.chatStorage.CreateChatMessage(context.Background(), message.Message, msg); err != nil {
						log.Printf("failed to create chat message in db: %v", err)
					}
				}

				return
			}

			contentStack += response.Choices[0].Delta.Content

			msg := entity.MessageCompletionStream{
				Event: entity.StreamEventMessageComletion,
				Data: entity.StreamData{
					Content: contentStack,
				},
				Topic:  topic,
				Status: entity.MessageStatusInProccess,
			}

			if err := stream.MessageCompletionStream.Write(msg, topic); err != nil {
				log.Printf("%v - ending stream for '%s'\n", err, topic)

				msg.Status = entity.MessageStatusInterrupted
				if contentStack != "" {
					if err := llm.chatStorage.CreateChatMessage(context.Background(), message.Message, msg); err != nil {
						log.Printf("failed to create chat message in db: %v", err)
					}
				}

				return
			}

			//some delay for visibility
			time.Sleep(time.Millisecond * 700)
		}
	}
}
