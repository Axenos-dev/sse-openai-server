// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ChatDbsColumns holds the columns for the "chat_dbs" table.
	ChatDbsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "topic", Type: field.TypeString},
		{Name: "message", Type: field.TypeString},
		{Name: "response", Type: field.TypeString},
		{Name: "status", Type: field.TypeInt},
		{Name: "timestamp", Type: field.TypeInt64},
	}
	// ChatDbsTable holds the schema information for the "chat_dbs" table.
	ChatDbsTable = &schema.Table{
		Name:       "chat_dbs",
		Columns:    ChatDbsColumns,
		PrimaryKey: []*schema.Column{ChatDbsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ChatDbsTable,
	}
)

func init() {
}
