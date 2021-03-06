// Code generated by entc, DO NOT EDIT.

package ent

import (
	"ShoppingBasket/ent/basket"
	"ShoppingBasket/ent/basketline"
	"ShoppingBasket/ent/event"
	"ShoppingBasket/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BasketLineUpdate is the builder for updating BasketLine entities.
type BasketLineUpdate struct {
	config
	hooks    []Hook
	mutation *BasketLineMutation
}

// Where appends a list predicates to the BasketLineUpdate builder.
func (blu *BasketLineUpdate) Where(ps ...predicate.BasketLine) *BasketLineUpdate {
	blu.mutation.Where(ps...)
	return blu
}

// SetTicketAmount sets the "TicketAmount" field.
func (blu *BasketLineUpdate) SetTicketAmount(u uint) *BasketLineUpdate {
	blu.mutation.ResetTicketAmount()
	blu.mutation.SetTicketAmount(u)
	return blu
}

// AddTicketAmount adds u to the "TicketAmount" field.
func (blu *BasketLineUpdate) AddTicketAmount(u int) *BasketLineUpdate {
	blu.mutation.AddTicketAmount(u)
	return blu
}

// SetPrice sets the "Price" field.
func (blu *BasketLineUpdate) SetPrice(f float32) *BasketLineUpdate {
	blu.mutation.ResetPrice()
	blu.mutation.SetPrice(f)
	return blu
}

// AddPrice adds f to the "Price" field.
func (blu *BasketLineUpdate) AddPrice(f float32) *BasketLineUpdate {
	blu.mutation.AddPrice(f)
	return blu
}

// SetEventID sets the "Event" edge to the Event entity by ID.
func (blu *BasketLineUpdate) SetEventID(id uuid.UUID) *BasketLineUpdate {
	blu.mutation.SetEventID(id)
	return blu
}

// SetNillableEventID sets the "Event" edge to the Event entity by ID if the given value is not nil.
func (blu *BasketLineUpdate) SetNillableEventID(id *uuid.UUID) *BasketLineUpdate {
	if id != nil {
		blu = blu.SetEventID(*id)
	}
	return blu
}

// SetEvent sets the "Event" edge to the Event entity.
func (blu *BasketLineUpdate) SetEvent(e *Event) *BasketLineUpdate {
	return blu.SetEventID(e.ID)
}

// SetBasketID sets the "Basket" edge to the Basket entity by ID.
func (blu *BasketLineUpdate) SetBasketID(id uuid.UUID) *BasketLineUpdate {
	blu.mutation.SetBasketID(id)
	return blu
}

// SetNillableBasketID sets the "Basket" edge to the Basket entity by ID if the given value is not nil.
func (blu *BasketLineUpdate) SetNillableBasketID(id *uuid.UUID) *BasketLineUpdate {
	if id != nil {
		blu = blu.SetBasketID(*id)
	}
	return blu
}

// SetBasket sets the "Basket" edge to the Basket entity.
func (blu *BasketLineUpdate) SetBasket(b *Basket) *BasketLineUpdate {
	return blu.SetBasketID(b.ID)
}

// Mutation returns the BasketLineMutation object of the builder.
func (blu *BasketLineUpdate) Mutation() *BasketLineMutation {
	return blu.mutation
}

// ClearEvent clears the "Event" edge to the Event entity.
func (blu *BasketLineUpdate) ClearEvent() *BasketLineUpdate {
	blu.mutation.ClearEvent()
	return blu
}

// ClearBasket clears the "Basket" edge to the Basket entity.
func (blu *BasketLineUpdate) ClearBasket() *BasketLineUpdate {
	blu.mutation.ClearBasket()
	return blu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (blu *BasketLineUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(blu.hooks) == 0 {
		affected, err = blu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BasketLineMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			blu.mutation = mutation
			affected, err = blu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(blu.hooks) - 1; i >= 0; i-- {
			if blu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = blu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, blu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (blu *BasketLineUpdate) SaveX(ctx context.Context) int {
	affected, err := blu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (blu *BasketLineUpdate) Exec(ctx context.Context) error {
	_, err := blu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (blu *BasketLineUpdate) ExecX(ctx context.Context) {
	if err := blu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (blu *BasketLineUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   basketline.Table,
			Columns: basketline.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: basketline.FieldID,
			},
		},
	}
	if ps := blu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := blu.mutation.TicketAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: basketline.FieldTicketAmount,
		})
	}
	if value, ok := blu.mutation.AddedTicketAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: basketline.FieldTicketAmount,
		})
	}
	if value, ok := blu.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: basketline.FieldPrice,
		})
	}
	if value, ok := blu.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: basketline.FieldPrice,
		})
	}
	if blu.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   basketline.EventTable,
			Columns: []string{basketline.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: event.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := blu.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   basketline.EventTable,
			Columns: []string{basketline.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: event.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if blu.mutation.BasketCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   basketline.BasketTable,
			Columns: []string{basketline.BasketColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: basket.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := blu.mutation.BasketIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   basketline.BasketTable,
			Columns: []string{basketline.BasketColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: basket.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, blu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{basketline.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// BasketLineUpdateOne is the builder for updating a single BasketLine entity.
type BasketLineUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BasketLineMutation
}

// SetTicketAmount sets the "TicketAmount" field.
func (bluo *BasketLineUpdateOne) SetTicketAmount(u uint) *BasketLineUpdateOne {
	bluo.mutation.ResetTicketAmount()
	bluo.mutation.SetTicketAmount(u)
	return bluo
}

// AddTicketAmount adds u to the "TicketAmount" field.
func (bluo *BasketLineUpdateOne) AddTicketAmount(u int) *BasketLineUpdateOne {
	bluo.mutation.AddTicketAmount(u)
	return bluo
}

// SetPrice sets the "Price" field.
func (bluo *BasketLineUpdateOne) SetPrice(f float32) *BasketLineUpdateOne {
	bluo.mutation.ResetPrice()
	bluo.mutation.SetPrice(f)
	return bluo
}

// AddPrice adds f to the "Price" field.
func (bluo *BasketLineUpdateOne) AddPrice(f float32) *BasketLineUpdateOne {
	bluo.mutation.AddPrice(f)
	return bluo
}

// SetEventID sets the "Event" edge to the Event entity by ID.
func (bluo *BasketLineUpdateOne) SetEventID(id uuid.UUID) *BasketLineUpdateOne {
	bluo.mutation.SetEventID(id)
	return bluo
}

// SetNillableEventID sets the "Event" edge to the Event entity by ID if the given value is not nil.
func (bluo *BasketLineUpdateOne) SetNillableEventID(id *uuid.UUID) *BasketLineUpdateOne {
	if id != nil {
		bluo = bluo.SetEventID(*id)
	}
	return bluo
}

// SetEvent sets the "Event" edge to the Event entity.
func (bluo *BasketLineUpdateOne) SetEvent(e *Event) *BasketLineUpdateOne {
	return bluo.SetEventID(e.ID)
}

// SetBasketID sets the "Basket" edge to the Basket entity by ID.
func (bluo *BasketLineUpdateOne) SetBasketID(id uuid.UUID) *BasketLineUpdateOne {
	bluo.mutation.SetBasketID(id)
	return bluo
}

// SetNillableBasketID sets the "Basket" edge to the Basket entity by ID if the given value is not nil.
func (bluo *BasketLineUpdateOne) SetNillableBasketID(id *uuid.UUID) *BasketLineUpdateOne {
	if id != nil {
		bluo = bluo.SetBasketID(*id)
	}
	return bluo
}

// SetBasket sets the "Basket" edge to the Basket entity.
func (bluo *BasketLineUpdateOne) SetBasket(b *Basket) *BasketLineUpdateOne {
	return bluo.SetBasketID(b.ID)
}

// Mutation returns the BasketLineMutation object of the builder.
func (bluo *BasketLineUpdateOne) Mutation() *BasketLineMutation {
	return bluo.mutation
}

// ClearEvent clears the "Event" edge to the Event entity.
func (bluo *BasketLineUpdateOne) ClearEvent() *BasketLineUpdateOne {
	bluo.mutation.ClearEvent()
	return bluo
}

// ClearBasket clears the "Basket" edge to the Basket entity.
func (bluo *BasketLineUpdateOne) ClearBasket() *BasketLineUpdateOne {
	bluo.mutation.ClearBasket()
	return bluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bluo *BasketLineUpdateOne) Select(field string, fields ...string) *BasketLineUpdateOne {
	bluo.fields = append([]string{field}, fields...)
	return bluo
}

// Save executes the query and returns the updated BasketLine entity.
func (bluo *BasketLineUpdateOne) Save(ctx context.Context) (*BasketLine, error) {
	var (
		err  error
		node *BasketLine
	)
	if len(bluo.hooks) == 0 {
		node, err = bluo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BasketLineMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bluo.mutation = mutation
			node, err = bluo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bluo.hooks) - 1; i >= 0; i-- {
			if bluo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bluo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bluo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (bluo *BasketLineUpdateOne) SaveX(ctx context.Context) *BasketLine {
	node, err := bluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bluo *BasketLineUpdateOne) Exec(ctx context.Context) error {
	_, err := bluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bluo *BasketLineUpdateOne) ExecX(ctx context.Context) {
	if err := bluo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bluo *BasketLineUpdateOne) sqlSave(ctx context.Context) (_node *BasketLine, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   basketline.Table,
			Columns: basketline.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: basketline.FieldID,
			},
		},
	}
	id, ok := bluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "BasketLine.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, basketline.FieldID)
		for _, f := range fields {
			if !basketline.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != basketline.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bluo.mutation.TicketAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: basketline.FieldTicketAmount,
		})
	}
	if value, ok := bluo.mutation.AddedTicketAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: basketline.FieldTicketAmount,
		})
	}
	if value, ok := bluo.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: basketline.FieldPrice,
		})
	}
	if value, ok := bluo.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: basketline.FieldPrice,
		})
	}
	if bluo.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   basketline.EventTable,
			Columns: []string{basketline.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: event.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bluo.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   basketline.EventTable,
			Columns: []string{basketline.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: event.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bluo.mutation.BasketCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   basketline.BasketTable,
			Columns: []string{basketline.BasketColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: basket.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bluo.mutation.BasketIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   basketline.BasketTable,
			Columns: []string{basketline.BasketColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: basket.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &BasketLine{config: bluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{basketline.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
