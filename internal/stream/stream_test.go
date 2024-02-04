package stream

import (
	"testing"
	"time"

	"github.com/Axenos-dev/sse-openai-server/internal/entity"
)

func TestStreamInitialization(t *testing.T) {
	topic := "1"
	err := MessageCompletionStream.InitNewStream(topic)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	defer MessageCompletionStream.CloseStream(topic)

	err = MessageCompletionStream.InitNewStream(topic)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestStreamWriteAndRead(t *testing.T) {
	topic := "2"
	err := MessageCompletionStream.InitNewStream(topic)
	if err != nil {
		t.Fatalf("Initialization error: %v", err)
	}

	defer MessageCompletionStream.CloseStream(topic)

	message := entity.MessageCompletionStream{
		Event: entity.StreamEventMessageComletion,
		Topic: topic,
		Data: entity.StreamData{
			Content: "Hello world!",
		},
	}

	go func() {
		time.Sleep(100 * time.Millisecond)
		MessageCompletionStream.Write(message, topic)
	}()

	select {
	case receivedMessage := <-MessageCompletionStream.Chan(topic):
		if receivedMessage != message {
			t.Errorf("Expected message %v, got %v", message, receivedMessage)
		}
	case <-time.After(500 * time.Millisecond):
		t.Error("Timeout waiting for message")
	}
}

func TestEndStream(t *testing.T) {
	topic := "3"
	err := MessageCompletionStream.InitNewStream(topic)
	if err != nil {
		t.Fatalf("Initialization error: %v", err)
	}
	defer MessageCompletionStream.CloseStream(topic)

	MessageCompletionStream.Write(entity.MessageCompletionStream{
		Event: entity.StreamEventMessageComletion,
		Topic: topic,
		Data: entity.StreamData{
			Content: "Hello world!",
		},
	}, topic)

	MessageCompletionStream.EndStream(topic)

	select {
	case <-MessageCompletionStream.EndChan(topic):
		// As expected
	case <-time.After(500 * time.Millisecond):
		t.Error("Timeout waiting for stream to end")
	}
}

func TestCloseStream(t *testing.T) {
	topic := "4"
	err := MessageCompletionStream.InitNewStream(topic)
	if err != nil {
		t.Fatalf("Initialization error: %v", err)
	}
	defer MessageCompletionStream.CloseStream(topic)

	MessageCompletionStream.CloseStream(topic)

	message := entity.MessageCompletionStream{
		Event:   entity.StreamEventMessageComletion,
		Message: "Test message",
	}
	err = MessageCompletionStream.Write(message, topic)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestDoesStreamExist(t *testing.T) {
	topic := "5"
	exists := MessageCompletionStream.DoesStreamExist(topic)
	if exists {
		t.Error("Expected stream not to exist, but it does")
	}
	defer MessageCompletionStream.CloseStream(topic)

	err := MessageCompletionStream.InitNewStream(topic)
	if err != nil {
		t.Fatalf("Initialization error: %v", err)
	}

	exists = MessageCompletionStream.DoesStreamExist(topic)
	if !exists {
		t.Error("Expected stream to exist, but it does not")
	}
}
