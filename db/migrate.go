package db

import "context"

func (db *DB) Migrate() error {
	return db.client.Schema.Create(context.Background())
}
