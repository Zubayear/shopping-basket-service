// Code generated by entc, DO NOT EDIT.

package ent

import (
	"ShoppingBasket/ent/basket"
	"ShoppingBasket/ent/basketline"
	"ShoppingBasket/ent/event"
	"ShoppingBasket/ent/predicate"
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BasketLineQuery is the builder for querying BasketLine entities.
type BasketLineQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.BasketLine
	// eager-loading edges.
	withEvent  *EventQuery
	withBasket *BasketQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BasketLineQuery builder.
func (blq *BasketLineQuery) Where(ps ...predicate.BasketLine) *BasketLineQuery {
	blq.predicates = append(blq.predicates, ps...)
	return blq
}

// Limit adds a limit step to the query.
func (blq *BasketLineQuery) Limit(limit int) *BasketLineQuery {
	blq.limit = &limit
	return blq
}

// Offset adds an offset step to the query.
func (blq *BasketLineQuery) Offset(offset int) *BasketLineQuery {
	blq.offset = &offset
	return blq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (blq *BasketLineQuery) Unique(unique bool) *BasketLineQuery {
	blq.unique = &unique
	return blq
}

// Order adds an order step to the query.
func (blq *BasketLineQuery) Order(o ...OrderFunc) *BasketLineQuery {
	blq.order = append(blq.order, o...)
	return blq
}

// QueryEvent chains the current query on the "Event" edge.
func (blq *BasketLineQuery) QueryEvent() *EventQuery {
	query := &EventQuery{config: blq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := blq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := blq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(basketline.Table, basketline.FieldID, selector),
			sqlgraph.To(event.Table, event.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, basketline.EventTable, basketline.EventColumn),
		)
		fromU = sqlgraph.SetNeighbors(blq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBasket chains the current query on the "Basket" edge.
func (blq *BasketLineQuery) QueryBasket() *BasketQuery {
	query := &BasketQuery{config: blq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := blq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := blq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(basketline.Table, basketline.FieldID, selector),
			sqlgraph.To(basket.Table, basket.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, basketline.BasketTable, basketline.BasketColumn),
		)
		fromU = sqlgraph.SetNeighbors(blq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BasketLine entity from the query.
// Returns a *NotFoundError when no BasketLine was found.
func (blq *BasketLineQuery) First(ctx context.Context) (*BasketLine, error) {
	nodes, err := blq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{basketline.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (blq *BasketLineQuery) FirstX(ctx context.Context) *BasketLine {
	node, err := blq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BasketLine ID from the query.
// Returns a *NotFoundError when no BasketLine ID was found.
func (blq *BasketLineQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = blq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{basketline.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (blq *BasketLineQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := blq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BasketLine entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BasketLine entity is found.
// Returns a *NotFoundError when no BasketLine entities are found.
func (blq *BasketLineQuery) Only(ctx context.Context) (*BasketLine, error) {
	nodes, err := blq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{basketline.Label}
	default:
		return nil, &NotSingularError{basketline.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (blq *BasketLineQuery) OnlyX(ctx context.Context) *BasketLine {
	node, err := blq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BasketLine ID in the query.
// Returns a *NotSingularError when more than one BasketLine ID is found.
// Returns a *NotFoundError when no entities are found.
func (blq *BasketLineQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = blq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{basketline.Label}
	default:
		err = &NotSingularError{basketline.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (blq *BasketLineQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := blq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BasketLines.
func (blq *BasketLineQuery) All(ctx context.Context) ([]*BasketLine, error) {
	if err := blq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return blq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (blq *BasketLineQuery) AllX(ctx context.Context) []*BasketLine {
	nodes, err := blq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BasketLine IDs.
func (blq *BasketLineQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := blq.Select(basketline.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (blq *BasketLineQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := blq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (blq *BasketLineQuery) Count(ctx context.Context) (int, error) {
	if err := blq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return blq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (blq *BasketLineQuery) CountX(ctx context.Context) int {
	count, err := blq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (blq *BasketLineQuery) Exist(ctx context.Context) (bool, error) {
	if err := blq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return blq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (blq *BasketLineQuery) ExistX(ctx context.Context) bool {
	exist, err := blq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BasketLineQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (blq *BasketLineQuery) Clone() *BasketLineQuery {
	if blq == nil {
		return nil
	}
	return &BasketLineQuery{
		config:     blq.config,
		limit:      blq.limit,
		offset:     blq.offset,
		order:      append([]OrderFunc{}, blq.order...),
		predicates: append([]predicate.BasketLine{}, blq.predicates...),
		withEvent:  blq.withEvent.Clone(),
		withBasket: blq.withBasket.Clone(),
		// clone intermediate query.
		sql:    blq.sql.Clone(),
		path:   blq.path,
		unique: blq.unique,
	}
}

// WithEvent tells the query-builder to eager-load the nodes that are connected to
// the "Event" edge. The optional arguments are used to configure the query builder of the edge.
func (blq *BasketLineQuery) WithEvent(opts ...func(*EventQuery)) *BasketLineQuery {
	query := &EventQuery{config: blq.config}
	for _, opt := range opts {
		opt(query)
	}
	blq.withEvent = query
	return blq
}

// WithBasket tells the query-builder to eager-load the nodes that are connected to
// the "Basket" edge. The optional arguments are used to configure the query builder of the edge.
func (blq *BasketLineQuery) WithBasket(opts ...func(*BasketQuery)) *BasketLineQuery {
	query := &BasketQuery{config: blq.config}
	for _, opt := range opts {
		opt(query)
	}
	blq.withBasket = query
	return blq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		TicketAmount uint `json:"TicketAmount,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BasketLine.Query().
//		GroupBy(basketline.FieldTicketAmount).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (blq *BasketLineQuery) GroupBy(field string, fields ...string) *BasketLineGroupBy {
	group := &BasketLineGroupBy{config: blq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := blq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return blq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		TicketAmount uint `json:"TicketAmount,omitempty"`
//	}
//
//	client.BasketLine.Query().
//		Select(basketline.FieldTicketAmount).
//		Scan(ctx, &v)
//
func (blq *BasketLineQuery) Select(fields ...string) *BasketLineSelect {
	blq.fields = append(blq.fields, fields...)
	return &BasketLineSelect{BasketLineQuery: blq}
}

func (blq *BasketLineQuery) prepareQuery(ctx context.Context) error {
	for _, f := range blq.fields {
		if !basketline.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if blq.path != nil {
		prev, err := blq.path(ctx)
		if err != nil {
			return err
		}
		blq.sql = prev
	}
	return nil
}

func (blq *BasketLineQuery) sqlAll(ctx context.Context) ([]*BasketLine, error) {
	var (
		nodes       = []*BasketLine{}
		withFKs     = blq.withFKs
		_spec       = blq.querySpec()
		loadedTypes = [2]bool{
			blq.withEvent != nil,
			blq.withBasket != nil,
		}
	)
	if blq.withBasket != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, basketline.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &BasketLine{config: blq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, blq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := blq.withEvent; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*BasketLine)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Event(func(s *sql.Selector) {
			s.Where(sql.InValues(basketline.EventColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.basket_line_event
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "basket_line_event" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "basket_line_event" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Event = n
		}
	}

	if query := blq.withBasket; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*BasketLine)
		for i := range nodes {
			if nodes[i].basket_basket_line == nil {
				continue
			}
			fk := *nodes[i].basket_basket_line
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(basket.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "basket_basket_line" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Basket = n
			}
		}
	}

	return nodes, nil
}

func (blq *BasketLineQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := blq.querySpec()
	_spec.Node.Columns = blq.fields
	if len(blq.fields) > 0 {
		_spec.Unique = blq.unique != nil && *blq.unique
	}
	return sqlgraph.CountNodes(ctx, blq.driver, _spec)
}

func (blq *BasketLineQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := blq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (blq *BasketLineQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   basketline.Table,
			Columns: basketline.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: basketline.FieldID,
			},
		},
		From:   blq.sql,
		Unique: true,
	}
	if unique := blq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := blq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, basketline.FieldID)
		for i := range fields {
			if fields[i] != basketline.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := blq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := blq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := blq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := blq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (blq *BasketLineQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(blq.driver.Dialect())
	t1 := builder.Table(basketline.Table)
	columns := blq.fields
	if len(columns) == 0 {
		columns = basketline.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if blq.sql != nil {
		selector = blq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if blq.unique != nil && *blq.unique {
		selector.Distinct()
	}
	for _, p := range blq.predicates {
		p(selector)
	}
	for _, p := range blq.order {
		p(selector)
	}
	if offset := blq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := blq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BasketLineGroupBy is the group-by builder for BasketLine entities.
type BasketLineGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (blgb *BasketLineGroupBy) Aggregate(fns ...AggregateFunc) *BasketLineGroupBy {
	blgb.fns = append(blgb.fns, fns...)
	return blgb
}

// Scan applies the group-by query and scans the result into the given value.
func (blgb *BasketLineGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := blgb.path(ctx)
	if err != nil {
		return err
	}
	blgb.sql = query
	return blgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (blgb *BasketLineGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := blgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (blgb *BasketLineGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(blgb.fields) > 1 {
		return nil, errors.New("ent: BasketLineGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := blgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (blgb *BasketLineGroupBy) StringsX(ctx context.Context) []string {
	v, err := blgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (blgb *BasketLineGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = blgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{basketline.Label}
	default:
		err = fmt.Errorf("ent: BasketLineGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (blgb *BasketLineGroupBy) StringX(ctx context.Context) string {
	v, err := blgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (blgb *BasketLineGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(blgb.fields) > 1 {
		return nil, errors.New("ent: BasketLineGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := blgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (blgb *BasketLineGroupBy) IntsX(ctx context.Context) []int {
	v, err := blgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (blgb *BasketLineGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = blgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{basketline.Label}
	default:
		err = fmt.Errorf("ent: BasketLineGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (blgb *BasketLineGroupBy) IntX(ctx context.Context) int {
	v, err := blgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (blgb *BasketLineGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(blgb.fields) > 1 {
		return nil, errors.New("ent: BasketLineGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := blgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (blgb *BasketLineGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := blgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (blgb *BasketLineGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = blgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{basketline.Label}
	default:
		err = fmt.Errorf("ent: BasketLineGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (blgb *BasketLineGroupBy) Float64X(ctx context.Context) float64 {
	v, err := blgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (blgb *BasketLineGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(blgb.fields) > 1 {
		return nil, errors.New("ent: BasketLineGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := blgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (blgb *BasketLineGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := blgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (blgb *BasketLineGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = blgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{basketline.Label}
	default:
		err = fmt.Errorf("ent: BasketLineGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (blgb *BasketLineGroupBy) BoolX(ctx context.Context) bool {
	v, err := blgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (blgb *BasketLineGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range blgb.fields {
		if !basketline.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := blgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := blgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (blgb *BasketLineGroupBy) sqlQuery() *sql.Selector {
	selector := blgb.sql.Select()
	aggregation := make([]string, 0, len(blgb.fns))
	for _, fn := range blgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(blgb.fields)+len(blgb.fns))
		for _, f := range blgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(blgb.fields...)...)
}

// BasketLineSelect is the builder for selecting fields of BasketLine entities.
type BasketLineSelect struct {
	*BasketLineQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bls *BasketLineSelect) Scan(ctx context.Context, v interface{}) error {
	if err := bls.prepareQuery(ctx); err != nil {
		return err
	}
	bls.sql = bls.BasketLineQuery.sqlQuery(ctx)
	return bls.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (bls *BasketLineSelect) ScanX(ctx context.Context, v interface{}) {
	if err := bls.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (bls *BasketLineSelect) Strings(ctx context.Context) ([]string, error) {
	if len(bls.fields) > 1 {
		return nil, errors.New("ent: BasketLineSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := bls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (bls *BasketLineSelect) StringsX(ctx context.Context) []string {
	v, err := bls.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (bls *BasketLineSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = bls.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{basketline.Label}
	default:
		err = fmt.Errorf("ent: BasketLineSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (bls *BasketLineSelect) StringX(ctx context.Context) string {
	v, err := bls.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (bls *BasketLineSelect) Ints(ctx context.Context) ([]int, error) {
	if len(bls.fields) > 1 {
		return nil, errors.New("ent: BasketLineSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := bls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (bls *BasketLineSelect) IntsX(ctx context.Context) []int {
	v, err := bls.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (bls *BasketLineSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = bls.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{basketline.Label}
	default:
		err = fmt.Errorf("ent: BasketLineSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (bls *BasketLineSelect) IntX(ctx context.Context) int {
	v, err := bls.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (bls *BasketLineSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(bls.fields) > 1 {
		return nil, errors.New("ent: BasketLineSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := bls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (bls *BasketLineSelect) Float64sX(ctx context.Context) []float64 {
	v, err := bls.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (bls *BasketLineSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = bls.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{basketline.Label}
	default:
		err = fmt.Errorf("ent: BasketLineSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (bls *BasketLineSelect) Float64X(ctx context.Context) float64 {
	v, err := bls.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (bls *BasketLineSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(bls.fields) > 1 {
		return nil, errors.New("ent: BasketLineSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := bls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (bls *BasketLineSelect) BoolsX(ctx context.Context) []bool {
	v, err := bls.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (bls *BasketLineSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = bls.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{basketline.Label}
	default:
		err = fmt.Errorf("ent: BasketLineSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (bls *BasketLineSelect) BoolX(ctx context.Context) bool {
	v, err := bls.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bls *BasketLineSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bls.sql.Query()
	if err := bls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
