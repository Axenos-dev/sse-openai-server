// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Axenos-dev/sse-openai-server/ent/chatdb"
	"github.com/Axenos-dev/sse-openai-server/ent/predicate"
)

// ChatDBDelete is the builder for deleting a ChatDB entity.
type ChatDBDelete struct {
	config
	hooks    []Hook
	mutation *ChatDBMutation
}

// Where appends a list predicates to the ChatDBDelete builder.
func (cdd *ChatDBDelete) Where(ps ...predicate.ChatDB) *ChatDBDelete {
	cdd.mutation.Where(ps...)
	return cdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cdd *ChatDBDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cdd.sqlExec, cdd.mutation, cdd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cdd *ChatDBDelete) ExecX(ctx context.Context) int {
	n, err := cdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cdd *ChatDBDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(chatdb.Table, sqlgraph.NewFieldSpec(chatdb.FieldID, field.TypeInt))
	if ps := cdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cdd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cdd.mutation.done = true
	return affected, err
}

// ChatDBDeleteOne is the builder for deleting a single ChatDB entity.
type ChatDBDeleteOne struct {
	cdd *ChatDBDelete
}

// Where appends a list predicates to the ChatDBDelete builder.
func (cddo *ChatDBDeleteOne) Where(ps ...predicate.ChatDB) *ChatDBDeleteOne {
	cddo.cdd.mutation.Where(ps...)
	return cddo
}

// Exec executes the deletion query.
func (cddo *ChatDBDeleteOne) Exec(ctx context.Context) error {
	n, err := cddo.cdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{chatdb.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cddo *ChatDBDeleteOne) ExecX(ctx context.Context) {
	if err := cddo.Exec(ctx); err != nil {
		panic(err)
	}
}
