package storage

import (
	"context"

	"github.com/Axenos-dev/sse-openai-server/db"
	"github.com/Axenos-dev/sse-openai-server/ent"
	"github.com/Axenos-dev/sse-openai-server/ent/chatdb"
	"github.com/Axenos-dev/sse-openai-server/internal/entity"
)

type ChatStorage struct {
	db.DB
}

func NewChatStorage(db db.DB) ChatStorage {
	return ChatStorage{db}
}

func (s ChatStorage) CreateChatMessage(ctx context.Context, message string, msg entity.MessageCompletionStream) error {
	_, err := s.DB.Client().ChatDB.
		Create().
		SetMessage(message).
		SetResponse(msg.Data.Content).
		SetTopic(msg.Topic).
		SetStatus(int(msg.Status)).
		Save(ctx)

	return err
}

func (s ChatStorage) QueryChats(ctx context.Context, topic string) ([]*ent.ChatDB, error) {
	return s.Client().ChatDB.Query().Where(chatdb.Topic(topic)).All(ctx)
}
