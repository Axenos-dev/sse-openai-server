package db

import (
	"fmt"

	"entgo.io/ent/dialect"
	"github.com/Axenos-dev/sse-openai-server/config"
	"github.com/Axenos-dev/sse-openai-server/ent"

	_ "github.com/lib/pq" // postgres driver
)

type DB struct {
	client *ent.Client
}

func NewDB(config config.PostgreSqlConfig) (db DB, err error) {
	db.client, err = ent.Open(dialect.Postgres, fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.User, config.Password, config.Host, config.Port, config.Database))
	return
}

func (db *DB) Client() *ent.Client {
	return db.client
}

func (db *DB) Close() {
	db.client.Close()
}
