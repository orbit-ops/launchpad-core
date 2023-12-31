// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/orbit-ops/launchpad-core/ent/access"
	"github.com/orbit-ops/launchpad-core/ent/actiontokens"
	"github.com/orbit-ops/launchpad-core/ent/predicate"
)

// ActionTokensUpdate is the builder for updating ActionTokens entities.
type ActionTokensUpdate struct {
	config
	hooks    []Hook
	mutation *ActionTokensMutation
}

// Where appends a list predicates to the ActionTokensUpdate builder.
func (atu *ActionTokensUpdate) Where(ps ...predicate.ActionTokens) *ActionTokensUpdate {
	atu.mutation.Where(ps...)
	return atu
}

// SetToken sets the "token" field.
func (atu *ActionTokensUpdate) SetToken(s string) *ActionTokensUpdate {
	atu.mutation.SetToken(s)
	return atu
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (atu *ActionTokensUpdate) SetNillableToken(s *string) *ActionTokensUpdate {
	if s != nil {
		atu.SetToken(*s)
	}
	return atu
}

// SetAccessID sets the "access_id" field.
func (atu *ActionTokensUpdate) SetAccessID(u uuid.UUID) *ActionTokensUpdate {
	atu.mutation.SetAccessID(u)
	return atu
}

// SetNillableAccessID sets the "access_id" field if the given value is not nil.
func (atu *ActionTokensUpdate) SetNillableAccessID(u *uuid.UUID) *ActionTokensUpdate {
	if u != nil {
		atu.SetAccessID(*u)
	}
	return atu
}

// SetAccessTokensID sets the "accessTokens" edge to the Access entity by ID.
func (atu *ActionTokensUpdate) SetAccessTokensID(id uuid.UUID) *ActionTokensUpdate {
	atu.mutation.SetAccessTokensID(id)
	return atu
}

// SetAccessTokens sets the "accessTokens" edge to the Access entity.
func (atu *ActionTokensUpdate) SetAccessTokens(a *Access) *ActionTokensUpdate {
	return atu.SetAccessTokensID(a.ID)
}

// Mutation returns the ActionTokensMutation object of the builder.
func (atu *ActionTokensUpdate) Mutation() *ActionTokensMutation {
	return atu.mutation
}

// ClearAccessTokens clears the "accessTokens" edge to the Access entity.
func (atu *ActionTokensUpdate) ClearAccessTokens() *ActionTokensUpdate {
	atu.mutation.ClearAccessTokens()
	return atu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (atu *ActionTokensUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, atu.sqlSave, atu.mutation, atu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (atu *ActionTokensUpdate) SaveX(ctx context.Context) int {
	affected, err := atu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (atu *ActionTokensUpdate) Exec(ctx context.Context) error {
	_, err := atu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atu *ActionTokensUpdate) ExecX(ctx context.Context) {
	if err := atu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atu *ActionTokensUpdate) check() error {
	if _, ok := atu.mutation.AccessTokensID(); atu.mutation.AccessTokensCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ActionTokens.accessTokens"`)
	}
	return nil
}

func (atu *ActionTokensUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := atu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(actiontokens.Table, actiontokens.Columns, sqlgraph.NewFieldSpec(actiontokens.FieldID, field.TypeUUID))
	if ps := atu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atu.mutation.Token(); ok {
		_spec.SetField(actiontokens.FieldToken, field.TypeString, value)
	}
	if atu.mutation.AccessTokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   actiontokens.AccessTokensTable,
			Columns: []string{actiontokens.AccessTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(access.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.AccessTokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   actiontokens.AccessTokensTable,
			Columns: []string{actiontokens.AccessTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(access.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, atu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{actiontokens.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	atu.mutation.done = true
	return n, nil
}

// ActionTokensUpdateOne is the builder for updating a single ActionTokens entity.
type ActionTokensUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ActionTokensMutation
}

// SetToken sets the "token" field.
func (atuo *ActionTokensUpdateOne) SetToken(s string) *ActionTokensUpdateOne {
	atuo.mutation.SetToken(s)
	return atuo
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (atuo *ActionTokensUpdateOne) SetNillableToken(s *string) *ActionTokensUpdateOne {
	if s != nil {
		atuo.SetToken(*s)
	}
	return atuo
}

// SetAccessID sets the "access_id" field.
func (atuo *ActionTokensUpdateOne) SetAccessID(u uuid.UUID) *ActionTokensUpdateOne {
	atuo.mutation.SetAccessID(u)
	return atuo
}

// SetNillableAccessID sets the "access_id" field if the given value is not nil.
func (atuo *ActionTokensUpdateOne) SetNillableAccessID(u *uuid.UUID) *ActionTokensUpdateOne {
	if u != nil {
		atuo.SetAccessID(*u)
	}
	return atuo
}

// SetAccessTokensID sets the "accessTokens" edge to the Access entity by ID.
func (atuo *ActionTokensUpdateOne) SetAccessTokensID(id uuid.UUID) *ActionTokensUpdateOne {
	atuo.mutation.SetAccessTokensID(id)
	return atuo
}

// SetAccessTokens sets the "accessTokens" edge to the Access entity.
func (atuo *ActionTokensUpdateOne) SetAccessTokens(a *Access) *ActionTokensUpdateOne {
	return atuo.SetAccessTokensID(a.ID)
}

// Mutation returns the ActionTokensMutation object of the builder.
func (atuo *ActionTokensUpdateOne) Mutation() *ActionTokensMutation {
	return atuo.mutation
}

// ClearAccessTokens clears the "accessTokens" edge to the Access entity.
func (atuo *ActionTokensUpdateOne) ClearAccessTokens() *ActionTokensUpdateOne {
	atuo.mutation.ClearAccessTokens()
	return atuo
}

// Where appends a list predicates to the ActionTokensUpdate builder.
func (atuo *ActionTokensUpdateOne) Where(ps ...predicate.ActionTokens) *ActionTokensUpdateOne {
	atuo.mutation.Where(ps...)
	return atuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (atuo *ActionTokensUpdateOne) Select(field string, fields ...string) *ActionTokensUpdateOne {
	atuo.fields = append([]string{field}, fields...)
	return atuo
}

// Save executes the query and returns the updated ActionTokens entity.
func (atuo *ActionTokensUpdateOne) Save(ctx context.Context) (*ActionTokens, error) {
	return withHooks(ctx, atuo.sqlSave, atuo.mutation, atuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (atuo *ActionTokensUpdateOne) SaveX(ctx context.Context) *ActionTokens {
	node, err := atuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (atuo *ActionTokensUpdateOne) Exec(ctx context.Context) error {
	_, err := atuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atuo *ActionTokensUpdateOne) ExecX(ctx context.Context) {
	if err := atuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atuo *ActionTokensUpdateOne) check() error {
	if _, ok := atuo.mutation.AccessTokensID(); atuo.mutation.AccessTokensCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ActionTokens.accessTokens"`)
	}
	return nil
}

func (atuo *ActionTokensUpdateOne) sqlSave(ctx context.Context) (_node *ActionTokens, err error) {
	if err := atuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(actiontokens.Table, actiontokens.Columns, sqlgraph.NewFieldSpec(actiontokens.FieldID, field.TypeUUID))
	id, ok := atuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ActionTokens.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := atuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, actiontokens.FieldID)
		for _, f := range fields {
			if !actiontokens.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != actiontokens.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := atuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atuo.mutation.Token(); ok {
		_spec.SetField(actiontokens.FieldToken, field.TypeString, value)
	}
	if atuo.mutation.AccessTokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   actiontokens.AccessTokensTable,
			Columns: []string{actiontokens.AccessTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(access.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.AccessTokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   actiontokens.AccessTokensTable,
			Columns: []string{actiontokens.AccessTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(access.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ActionTokens{config: atuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, atuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{actiontokens.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	atuo.mutation.done = true
	return _node, nil
}
