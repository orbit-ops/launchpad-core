// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/orbit-ops/launchpad-core/ent/access"
	"github.com/orbit-ops/launchpad-core/ent/actiontokens"
	"github.com/orbit-ops/launchpad-core/ent/predicate"
)

// ActionTokensQuery is the builder for querying ActionTokens entities.
type ActionTokensQuery struct {
	config
	ctx              *QueryContext
	order            []actiontokens.OrderOption
	inters           []Interceptor
	predicates       []predicate.ActionTokens
	withAccessTokens *AccessQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ActionTokensQuery builder.
func (atq *ActionTokensQuery) Where(ps ...predicate.ActionTokens) *ActionTokensQuery {
	atq.predicates = append(atq.predicates, ps...)
	return atq
}

// Limit the number of records to be returned by this query.
func (atq *ActionTokensQuery) Limit(limit int) *ActionTokensQuery {
	atq.ctx.Limit = &limit
	return atq
}

// Offset to start from.
func (atq *ActionTokensQuery) Offset(offset int) *ActionTokensQuery {
	atq.ctx.Offset = &offset
	return atq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (atq *ActionTokensQuery) Unique(unique bool) *ActionTokensQuery {
	atq.ctx.Unique = &unique
	return atq
}

// Order specifies how the records should be ordered.
func (atq *ActionTokensQuery) Order(o ...actiontokens.OrderOption) *ActionTokensQuery {
	atq.order = append(atq.order, o...)
	return atq
}

// QueryAccessTokens chains the current query on the "accessTokens" edge.
func (atq *ActionTokensQuery) QueryAccessTokens() *AccessQuery {
	query := (&AccessClient{config: atq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := atq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := atq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(actiontokens.Table, actiontokens.FieldID, selector),
			sqlgraph.To(access.Table, access.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, actiontokens.AccessTokensTable, actiontokens.AccessTokensColumn),
		)
		fromU = sqlgraph.SetNeighbors(atq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ActionTokens entity from the query.
// Returns a *NotFoundError when no ActionTokens was found.
func (atq *ActionTokensQuery) First(ctx context.Context) (*ActionTokens, error) {
	nodes, err := atq.Limit(1).All(setContextOp(ctx, atq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{actiontokens.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (atq *ActionTokensQuery) FirstX(ctx context.Context) *ActionTokens {
	node, err := atq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ActionTokens ID from the query.
// Returns a *NotFoundError when no ActionTokens ID was found.
func (atq *ActionTokensQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = atq.Limit(1).IDs(setContextOp(ctx, atq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{actiontokens.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (atq *ActionTokensQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := atq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ActionTokens entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ActionTokens entity is found.
// Returns a *NotFoundError when no ActionTokens entities are found.
func (atq *ActionTokensQuery) Only(ctx context.Context) (*ActionTokens, error) {
	nodes, err := atq.Limit(2).All(setContextOp(ctx, atq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{actiontokens.Label}
	default:
		return nil, &NotSingularError{actiontokens.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (atq *ActionTokensQuery) OnlyX(ctx context.Context) *ActionTokens {
	node, err := atq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ActionTokens ID in the query.
// Returns a *NotSingularError when more than one ActionTokens ID is found.
// Returns a *NotFoundError when no entities are found.
func (atq *ActionTokensQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = atq.Limit(2).IDs(setContextOp(ctx, atq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{actiontokens.Label}
	default:
		err = &NotSingularError{actiontokens.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (atq *ActionTokensQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := atq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ActionTokensSlice.
func (atq *ActionTokensQuery) All(ctx context.Context) ([]*ActionTokens, error) {
	ctx = setContextOp(ctx, atq.ctx, "All")
	if err := atq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ActionTokens, *ActionTokensQuery]()
	return withInterceptors[[]*ActionTokens](ctx, atq, qr, atq.inters)
}

// AllX is like All, but panics if an error occurs.
func (atq *ActionTokensQuery) AllX(ctx context.Context) []*ActionTokens {
	nodes, err := atq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ActionTokens IDs.
func (atq *ActionTokensQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if atq.ctx.Unique == nil && atq.path != nil {
		atq.Unique(true)
	}
	ctx = setContextOp(ctx, atq.ctx, "IDs")
	if err = atq.Select(actiontokens.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (atq *ActionTokensQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := atq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (atq *ActionTokensQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, atq.ctx, "Count")
	if err := atq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, atq, querierCount[*ActionTokensQuery](), atq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (atq *ActionTokensQuery) CountX(ctx context.Context) int {
	count, err := atq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (atq *ActionTokensQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, atq.ctx, "Exist")
	switch _, err := atq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (atq *ActionTokensQuery) ExistX(ctx context.Context) bool {
	exist, err := atq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ActionTokensQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (atq *ActionTokensQuery) Clone() *ActionTokensQuery {
	if atq == nil {
		return nil
	}
	return &ActionTokensQuery{
		config:           atq.config,
		ctx:              atq.ctx.Clone(),
		order:            append([]actiontokens.OrderOption{}, atq.order...),
		inters:           append([]Interceptor{}, atq.inters...),
		predicates:       append([]predicate.ActionTokens{}, atq.predicates...),
		withAccessTokens: atq.withAccessTokens.Clone(),
		// clone intermediate query.
		sql:  atq.sql.Clone(),
		path: atq.path,
	}
}

// WithAccessTokens tells the query-builder to eager-load the nodes that are connected to
// the "accessTokens" edge. The optional arguments are used to configure the query builder of the edge.
func (atq *ActionTokensQuery) WithAccessTokens(opts ...func(*AccessQuery)) *ActionTokensQuery {
	query := (&AccessClient{config: atq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	atq.withAccessTokens = query
	return atq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Action actiontokens.Action `json:"action,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ActionTokens.Query().
//		GroupBy(actiontokens.FieldAction).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (atq *ActionTokensQuery) GroupBy(field string, fields ...string) *ActionTokensGroupBy {
	atq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ActionTokensGroupBy{build: atq}
	grbuild.flds = &atq.ctx.Fields
	grbuild.label = actiontokens.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Action actiontokens.Action `json:"action,omitempty"`
//	}
//
//	client.ActionTokens.Query().
//		Select(actiontokens.FieldAction).
//		Scan(ctx, &v)
func (atq *ActionTokensQuery) Select(fields ...string) *ActionTokensSelect {
	atq.ctx.Fields = append(atq.ctx.Fields, fields...)
	sbuild := &ActionTokensSelect{ActionTokensQuery: atq}
	sbuild.label = actiontokens.Label
	sbuild.flds, sbuild.scan = &atq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ActionTokensSelect configured with the given aggregations.
func (atq *ActionTokensQuery) Aggregate(fns ...AggregateFunc) *ActionTokensSelect {
	return atq.Select().Aggregate(fns...)
}

func (atq *ActionTokensQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range atq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, atq); err != nil {
				return err
			}
		}
	}
	for _, f := range atq.ctx.Fields {
		if !actiontokens.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if atq.path != nil {
		prev, err := atq.path(ctx)
		if err != nil {
			return err
		}
		atq.sql = prev
	}
	return nil
}

func (atq *ActionTokensQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ActionTokens, error) {
	var (
		nodes       = []*ActionTokens{}
		_spec       = atq.querySpec()
		loadedTypes = [1]bool{
			atq.withAccessTokens != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ActionTokens).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ActionTokens{config: atq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, atq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := atq.withAccessTokens; query != nil {
		if err := atq.loadAccessTokens(ctx, query, nodes, nil,
			func(n *ActionTokens, e *Access) { n.Edges.AccessTokens = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (atq *ActionTokensQuery) loadAccessTokens(ctx context.Context, query *AccessQuery, nodes []*ActionTokens, init func(*ActionTokens), assign func(*ActionTokens, *Access)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ActionTokens)
	for i := range nodes {
		fk := nodes[i].AccessID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(access.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "access_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (atq *ActionTokensQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := atq.querySpec()
	_spec.Node.Columns = atq.ctx.Fields
	if len(atq.ctx.Fields) > 0 {
		_spec.Unique = atq.ctx.Unique != nil && *atq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, atq.driver, _spec)
}

func (atq *ActionTokensQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(actiontokens.Table, actiontokens.Columns, sqlgraph.NewFieldSpec(actiontokens.FieldID, field.TypeUUID))
	_spec.From = atq.sql
	if unique := atq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if atq.path != nil {
		_spec.Unique = true
	}
	if fields := atq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, actiontokens.FieldID)
		for i := range fields {
			if fields[i] != actiontokens.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if atq.withAccessTokens != nil {
			_spec.Node.AddColumnOnce(actiontokens.FieldAccessID)
		}
	}
	if ps := atq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := atq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := atq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := atq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (atq *ActionTokensQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(atq.driver.Dialect())
	t1 := builder.Table(actiontokens.Table)
	columns := atq.ctx.Fields
	if len(columns) == 0 {
		columns = actiontokens.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if atq.sql != nil {
		selector = atq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if atq.ctx.Unique != nil && *atq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range atq.predicates {
		p(selector)
	}
	for _, p := range atq.order {
		p(selector)
	}
	if offset := atq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := atq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ActionTokensGroupBy is the group-by builder for ActionTokens entities.
type ActionTokensGroupBy struct {
	selector
	build *ActionTokensQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (atgb *ActionTokensGroupBy) Aggregate(fns ...AggregateFunc) *ActionTokensGroupBy {
	atgb.fns = append(atgb.fns, fns...)
	return atgb
}

// Scan applies the selector query and scans the result into the given value.
func (atgb *ActionTokensGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, atgb.build.ctx, "GroupBy")
	if err := atgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ActionTokensQuery, *ActionTokensGroupBy](ctx, atgb.build, atgb, atgb.build.inters, v)
}

func (atgb *ActionTokensGroupBy) sqlScan(ctx context.Context, root *ActionTokensQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(atgb.fns))
	for _, fn := range atgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*atgb.flds)+len(atgb.fns))
		for _, f := range *atgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*atgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := atgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ActionTokensSelect is the builder for selecting fields of ActionTokens entities.
type ActionTokensSelect struct {
	*ActionTokensQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ats *ActionTokensSelect) Aggregate(fns ...AggregateFunc) *ActionTokensSelect {
	ats.fns = append(ats.fns, fns...)
	return ats
}

// Scan applies the selector query and scans the result into the given value.
func (ats *ActionTokensSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ats.ctx, "Select")
	if err := ats.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ActionTokensQuery, *ActionTokensSelect](ctx, ats.ActionTokensQuery, ats, ats.inters, v)
}

func (ats *ActionTokensSelect) sqlScan(ctx context.Context, root *ActionTokensQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ats.fns))
	for _, fn := range ats.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ats.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ats.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
