package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ChatDB holds the schema definition for the ChatDB entity.
type ChatDB struct {
	ent.Schema
}

// Fields of the ChatDB.
func (ChatDB) Fields() []ent.Field {
	return []ent.Field{
		field.String("topic"),
		field.String("message"),
		field.String("response"),
		field.Int("status").NonNegative(),
		field.Int64("timestamp").DefaultFunc(func() int64 {
			return time.Now().Unix()
		}),
	}
}

// Edges of the ChatDB.
func (ChatDB) Edges() []ent.Edge {
	return nil
}
