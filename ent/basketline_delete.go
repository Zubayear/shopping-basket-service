// Code generated by entc, DO NOT EDIT.

package ent

import (
	"ShoppingBasket/ent/basketline"
	"ShoppingBasket/ent/predicate"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BasketLineDelete is the builder for deleting a BasketLine entity.
type BasketLineDelete struct {
	config
	hooks    []Hook
	mutation *BasketLineMutation
}

// Where appends a list predicates to the BasketLineDelete builder.
func (bld *BasketLineDelete) Where(ps ...predicate.BasketLine) *BasketLineDelete {
	bld.mutation.Where(ps...)
	return bld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bld *BasketLineDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bld.hooks) == 0 {
		affected, err = bld.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BasketLineMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bld.mutation = mutation
			affected, err = bld.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bld.hooks) - 1; i >= 0; i-- {
			if bld.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bld.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bld.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (bld *BasketLineDelete) ExecX(ctx context.Context) int {
	n, err := bld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bld *BasketLineDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: basketline.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: basketline.FieldID,
			},
		},
	}
	if ps := bld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, bld.driver, _spec)
}

// BasketLineDeleteOne is the builder for deleting a single BasketLine entity.
type BasketLineDeleteOne struct {
	bld *BasketLineDelete
}

// Exec executes the deletion query.
func (bldo *BasketLineDeleteOne) Exec(ctx context.Context) error {
	n, err := bldo.bld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{basketline.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bldo *BasketLineDeleteOne) ExecX(ctx context.Context) {
	bldo.bld.ExecX(ctx)
}
