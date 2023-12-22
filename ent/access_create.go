// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/orbit-ops/launchpad-core/ent/access"
	"github.com/orbit-ops/launchpad-core/ent/actiontokens"
)

// AccessCreate is the builder for creating a Access entity.
type AccessCreate struct {
	config
	mutation *AccessMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetStartTime sets the "start_time" field.
func (ac *AccessCreate) SetStartTime(t time.Time) *AccessCreate {
	ac.mutation.SetStartTime(t)
	return ac
}

// SetApproved sets the "approved" field.
func (ac *AccessCreate) SetApproved(b bool) *AccessCreate {
	ac.mutation.SetApproved(b)
	return ac
}

// SetRolledBack sets the "rolled_back" field.
func (ac *AccessCreate) SetRolledBack(b bool) *AccessCreate {
	ac.mutation.SetRolledBack(b)
	return ac
}

// SetNillableRolledBack sets the "rolled_back" field if the given value is not nil.
func (ac *AccessCreate) SetNillableRolledBack(b *bool) *AccessCreate {
	if b != nil {
		ac.SetRolledBack(*b)
	}
	return ac
}

// SetRollbackTime sets the "rollback_time" field.
func (ac *AccessCreate) SetRollbackTime(t time.Time) *AccessCreate {
	ac.mutation.SetRollbackTime(t)
	return ac
}

// SetNillableRollbackTime sets the "rollback_time" field if the given value is not nil.
func (ac *AccessCreate) SetNillableRollbackTime(t *time.Time) *AccessCreate {
	if t != nil {
		ac.SetRollbackTime(*t)
	}
	return ac
}

// SetRollbackReason sets the "rollback_reason" field.
func (ac *AccessCreate) SetRollbackReason(s string) *AccessCreate {
	ac.mutation.SetRollbackReason(s)
	return ac
}

// SetNillableRollbackReason sets the "rollback_reason" field if the given value is not nil.
func (ac *AccessCreate) SetNillableRollbackReason(s *string) *AccessCreate {
	if s != nil {
		ac.SetRollbackReason(*s)
	}
	return ac
}

// SetEndTime sets the "end_time" field.
func (ac *AccessCreate) SetEndTime(t time.Time) *AccessCreate {
	ac.mutation.SetEndTime(t)
	return ac
}

// SetRequestID sets the "request_id" field.
func (ac *AccessCreate) SetRequestID(u uuid.UUID) *AccessCreate {
	ac.mutation.SetRequestID(u)
	return ac
}

// SetID sets the "id" field.
func (ac *AccessCreate) SetID(u uuid.UUID) *AccessCreate {
	ac.mutation.SetID(u)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *AccessCreate) SetNillableID(u *uuid.UUID) *AccessCreate {
	if u != nil {
		ac.SetID(*u)
	}
	return ac
}

// SetApprovalsID sets the "approvals" edge to the Access entity by ID.
func (ac *AccessCreate) SetApprovalsID(id uuid.UUID) *AccessCreate {
	ac.mutation.SetApprovalsID(id)
	return ac
}

// SetNillableApprovalsID sets the "approvals" edge to the Access entity by ID if the given value is not nil.
func (ac *AccessCreate) SetNillableApprovalsID(id *uuid.UUID) *AccessCreate {
	if id != nil {
		ac = ac.SetApprovalsID(*id)
	}
	return ac
}

// SetApprovals sets the "approvals" edge to the Access entity.
func (ac *AccessCreate) SetApprovals(a *Access) *AccessCreate {
	return ac.SetApprovalsID(a.ID)
}

// AddAccessTokenIDs adds the "accessTokens" edge to the ActionTokens entity by IDs.
func (ac *AccessCreate) AddAccessTokenIDs(ids ...uuid.UUID) *AccessCreate {
	ac.mutation.AddAccessTokenIDs(ids...)
	return ac
}

// AddAccessTokens adds the "accessTokens" edges to the ActionTokens entity.
func (ac *AccessCreate) AddAccessTokens(a ...*ActionTokens) *AccessCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ac.AddAccessTokenIDs(ids...)
}

// Mutation returns the AccessMutation object of the builder.
func (ac *AccessCreate) Mutation() *AccessMutation {
	return ac.mutation
}

// Save creates the Access in the database.
func (ac *AccessCreate) Save(ctx context.Context) (*Access, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AccessCreate) SaveX(ctx context.Context) *Access {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AccessCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AccessCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AccessCreate) defaults() {
	if _, ok := ac.mutation.RolledBack(); !ok {
		v := access.DefaultRolledBack
		ac.mutation.SetRolledBack(v)
	}
	if _, ok := ac.mutation.ID(); !ok {
		v := access.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AccessCreate) check() error {
	if _, ok := ac.mutation.StartTime(); !ok {
		return &ValidationError{Name: "start_time", err: errors.New(`ent: missing required field "Access.start_time"`)}
	}
	if _, ok := ac.mutation.Approved(); !ok {
		return &ValidationError{Name: "approved", err: errors.New(`ent: missing required field "Access.approved"`)}
	}
	if _, ok := ac.mutation.RolledBack(); !ok {
		return &ValidationError{Name: "rolled_back", err: errors.New(`ent: missing required field "Access.rolled_back"`)}
	}
	if _, ok := ac.mutation.EndTime(); !ok {
		return &ValidationError{Name: "end_time", err: errors.New(`ent: missing required field "Access.end_time"`)}
	}
	if _, ok := ac.mutation.RequestID(); !ok {
		return &ValidationError{Name: "request_id", err: errors.New(`ent: missing required field "Access.request_id"`)}
	}
	return nil
}

func (ac *AccessCreate) sqlSave(ctx context.Context) (*Access, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AccessCreate) createSpec() (*Access, *sqlgraph.CreateSpec) {
	var (
		_node = &Access{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(access.Table, sqlgraph.NewFieldSpec(access.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = ac.conflict
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ac.mutation.StartTime(); ok {
		_spec.SetField(access.FieldStartTime, field.TypeTime, value)
		_node.StartTime = value
	}
	if value, ok := ac.mutation.Approved(); ok {
		_spec.SetField(access.FieldApproved, field.TypeBool, value)
		_node.Approved = value
	}
	if value, ok := ac.mutation.RolledBack(); ok {
		_spec.SetField(access.FieldRolledBack, field.TypeBool, value)
		_node.RolledBack = value
	}
	if value, ok := ac.mutation.RollbackTime(); ok {
		_spec.SetField(access.FieldRollbackTime, field.TypeTime, value)
		_node.RollbackTime = value
	}
	if value, ok := ac.mutation.RollbackReason(); ok {
		_spec.SetField(access.FieldRollbackReason, field.TypeString, value)
		_node.RollbackReason = value
	}
	if value, ok := ac.mutation.EndTime(); ok {
		_spec.SetField(access.FieldEndTime, field.TypeTime, value)
		_node.EndTime = value
	}
	if value, ok := ac.mutation.RequestID(); ok {
		_spec.SetField(access.FieldRequestID, field.TypeUUID, value)
		_node.RequestID = value
	}
	if nodes := ac.mutation.ApprovalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   access.ApprovalsTable,
			Columns: []string{access.ApprovalsColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(access.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.access_approvals = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.AccessTokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   access.AccessTokensTable,
			Columns: []string{access.AccessTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(actiontokens.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Access.Create().
//		SetStartTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AccessUpsert) {
//			SetStartTime(v+v).
//		}).
//		Exec(ctx)
func (ac *AccessCreate) OnConflict(opts ...sql.ConflictOption) *AccessUpsertOne {
	ac.conflict = opts
	return &AccessUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Access.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ac *AccessCreate) OnConflictColumns(columns ...string) *AccessUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AccessUpsertOne{
		create: ac,
	}
}

type (
	// AccessUpsertOne is the builder for "upsert"-ing
	//  one Access node.
	AccessUpsertOne struct {
		create *AccessCreate
	}

	// AccessUpsert is the "OnConflict" setter.
	AccessUpsert struct {
		*sql.UpdateSet
	}
)

// SetRolledBack sets the "rolled_back" field.
func (u *AccessUpsert) SetRolledBack(v bool) *AccessUpsert {
	u.Set(access.FieldRolledBack, v)
	return u
}

// UpdateRolledBack sets the "rolled_back" field to the value that was provided on create.
func (u *AccessUpsert) UpdateRolledBack() *AccessUpsert {
	u.SetExcluded(access.FieldRolledBack)
	return u
}

// SetRollbackTime sets the "rollback_time" field.
func (u *AccessUpsert) SetRollbackTime(v time.Time) *AccessUpsert {
	u.Set(access.FieldRollbackTime, v)
	return u
}

// UpdateRollbackTime sets the "rollback_time" field to the value that was provided on create.
func (u *AccessUpsert) UpdateRollbackTime() *AccessUpsert {
	u.SetExcluded(access.FieldRollbackTime)
	return u
}

// ClearRollbackTime clears the value of the "rollback_time" field.
func (u *AccessUpsert) ClearRollbackTime() *AccessUpsert {
	u.SetNull(access.FieldRollbackTime)
	return u
}

// SetRollbackReason sets the "rollback_reason" field.
func (u *AccessUpsert) SetRollbackReason(v string) *AccessUpsert {
	u.Set(access.FieldRollbackReason, v)
	return u
}

// UpdateRollbackReason sets the "rollback_reason" field to the value that was provided on create.
func (u *AccessUpsert) UpdateRollbackReason() *AccessUpsert {
	u.SetExcluded(access.FieldRollbackReason)
	return u
}

// ClearRollbackReason clears the value of the "rollback_reason" field.
func (u *AccessUpsert) ClearRollbackReason() *AccessUpsert {
	u.SetNull(access.FieldRollbackReason)
	return u
}

// SetRequestID sets the "request_id" field.
func (u *AccessUpsert) SetRequestID(v uuid.UUID) *AccessUpsert {
	u.Set(access.FieldRequestID, v)
	return u
}

// UpdateRequestID sets the "request_id" field to the value that was provided on create.
func (u *AccessUpsert) UpdateRequestID() *AccessUpsert {
	u.SetExcluded(access.FieldRequestID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Access.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(access.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AccessUpsertOne) UpdateNewValues() *AccessUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(access.FieldID)
		}
		if _, exists := u.create.mutation.StartTime(); exists {
			s.SetIgnore(access.FieldStartTime)
		}
		if _, exists := u.create.mutation.Approved(); exists {
			s.SetIgnore(access.FieldApproved)
		}
		if _, exists := u.create.mutation.EndTime(); exists {
			s.SetIgnore(access.FieldEndTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Access.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AccessUpsertOne) Ignore() *AccessUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AccessUpsertOne) DoNothing() *AccessUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AccessCreate.OnConflict
// documentation for more info.
func (u *AccessUpsertOne) Update(set func(*AccessUpsert)) *AccessUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AccessUpsert{UpdateSet: update})
	}))
	return u
}

// SetRolledBack sets the "rolled_back" field.
func (u *AccessUpsertOne) SetRolledBack(v bool) *AccessUpsertOne {
	return u.Update(func(s *AccessUpsert) {
		s.SetRolledBack(v)
	})
}

// UpdateRolledBack sets the "rolled_back" field to the value that was provided on create.
func (u *AccessUpsertOne) UpdateRolledBack() *AccessUpsertOne {
	return u.Update(func(s *AccessUpsert) {
		s.UpdateRolledBack()
	})
}

// SetRollbackTime sets the "rollback_time" field.
func (u *AccessUpsertOne) SetRollbackTime(v time.Time) *AccessUpsertOne {
	return u.Update(func(s *AccessUpsert) {
		s.SetRollbackTime(v)
	})
}

// UpdateRollbackTime sets the "rollback_time" field to the value that was provided on create.
func (u *AccessUpsertOne) UpdateRollbackTime() *AccessUpsertOne {
	return u.Update(func(s *AccessUpsert) {
		s.UpdateRollbackTime()
	})
}

// ClearRollbackTime clears the value of the "rollback_time" field.
func (u *AccessUpsertOne) ClearRollbackTime() *AccessUpsertOne {
	return u.Update(func(s *AccessUpsert) {
		s.ClearRollbackTime()
	})
}

// SetRollbackReason sets the "rollback_reason" field.
func (u *AccessUpsertOne) SetRollbackReason(v string) *AccessUpsertOne {
	return u.Update(func(s *AccessUpsert) {
		s.SetRollbackReason(v)
	})
}

// UpdateRollbackReason sets the "rollback_reason" field to the value that was provided on create.
func (u *AccessUpsertOne) UpdateRollbackReason() *AccessUpsertOne {
	return u.Update(func(s *AccessUpsert) {
		s.UpdateRollbackReason()
	})
}

// ClearRollbackReason clears the value of the "rollback_reason" field.
func (u *AccessUpsertOne) ClearRollbackReason() *AccessUpsertOne {
	return u.Update(func(s *AccessUpsert) {
		s.ClearRollbackReason()
	})
}

// SetRequestID sets the "request_id" field.
func (u *AccessUpsertOne) SetRequestID(v uuid.UUID) *AccessUpsertOne {
	return u.Update(func(s *AccessUpsert) {
		s.SetRequestID(v)
	})
}

// UpdateRequestID sets the "request_id" field to the value that was provided on create.
func (u *AccessUpsertOne) UpdateRequestID() *AccessUpsertOne {
	return u.Update(func(s *AccessUpsert) {
		s.UpdateRequestID()
	})
}

// Exec executes the query.
func (u *AccessUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AccessCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AccessUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AccessUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: AccessUpsertOne.ID is not supported by MySQL driver. Use AccessUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AccessUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AccessCreateBulk is the builder for creating many Access entities in bulk.
type AccessCreateBulk struct {
	config
	err      error
	builders []*AccessCreate
	conflict []sql.ConflictOption
}

// Save creates the Access entities in the database.
func (acb *AccessCreateBulk) Save(ctx context.Context) ([]*Access, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Access, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AccessMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AccessCreateBulk) SaveX(ctx context.Context) []*Access {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AccessCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AccessCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Access.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AccessUpsert) {
//			SetStartTime(v+v).
//		}).
//		Exec(ctx)
func (acb *AccessCreateBulk) OnConflict(opts ...sql.ConflictOption) *AccessUpsertBulk {
	acb.conflict = opts
	return &AccessUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Access.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acb *AccessCreateBulk) OnConflictColumns(columns ...string) *AccessUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AccessUpsertBulk{
		create: acb,
	}
}

// AccessUpsertBulk is the builder for "upsert"-ing
// a bulk of Access nodes.
type AccessUpsertBulk struct {
	create *AccessCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Access.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(access.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AccessUpsertBulk) UpdateNewValues() *AccessUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(access.FieldID)
			}
			if _, exists := b.mutation.StartTime(); exists {
				s.SetIgnore(access.FieldStartTime)
			}
			if _, exists := b.mutation.Approved(); exists {
				s.SetIgnore(access.FieldApproved)
			}
			if _, exists := b.mutation.EndTime(); exists {
				s.SetIgnore(access.FieldEndTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Access.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AccessUpsertBulk) Ignore() *AccessUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AccessUpsertBulk) DoNothing() *AccessUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AccessCreateBulk.OnConflict
// documentation for more info.
func (u *AccessUpsertBulk) Update(set func(*AccessUpsert)) *AccessUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AccessUpsert{UpdateSet: update})
	}))
	return u
}

// SetRolledBack sets the "rolled_back" field.
func (u *AccessUpsertBulk) SetRolledBack(v bool) *AccessUpsertBulk {
	return u.Update(func(s *AccessUpsert) {
		s.SetRolledBack(v)
	})
}

// UpdateRolledBack sets the "rolled_back" field to the value that was provided on create.
func (u *AccessUpsertBulk) UpdateRolledBack() *AccessUpsertBulk {
	return u.Update(func(s *AccessUpsert) {
		s.UpdateRolledBack()
	})
}

// SetRollbackTime sets the "rollback_time" field.
func (u *AccessUpsertBulk) SetRollbackTime(v time.Time) *AccessUpsertBulk {
	return u.Update(func(s *AccessUpsert) {
		s.SetRollbackTime(v)
	})
}

// UpdateRollbackTime sets the "rollback_time" field to the value that was provided on create.
func (u *AccessUpsertBulk) UpdateRollbackTime() *AccessUpsertBulk {
	return u.Update(func(s *AccessUpsert) {
		s.UpdateRollbackTime()
	})
}

// ClearRollbackTime clears the value of the "rollback_time" field.
func (u *AccessUpsertBulk) ClearRollbackTime() *AccessUpsertBulk {
	return u.Update(func(s *AccessUpsert) {
		s.ClearRollbackTime()
	})
}

// SetRollbackReason sets the "rollback_reason" field.
func (u *AccessUpsertBulk) SetRollbackReason(v string) *AccessUpsertBulk {
	return u.Update(func(s *AccessUpsert) {
		s.SetRollbackReason(v)
	})
}

// UpdateRollbackReason sets the "rollback_reason" field to the value that was provided on create.
func (u *AccessUpsertBulk) UpdateRollbackReason() *AccessUpsertBulk {
	return u.Update(func(s *AccessUpsert) {
		s.UpdateRollbackReason()
	})
}

// ClearRollbackReason clears the value of the "rollback_reason" field.
func (u *AccessUpsertBulk) ClearRollbackReason() *AccessUpsertBulk {
	return u.Update(func(s *AccessUpsert) {
		s.ClearRollbackReason()
	})
}

// SetRequestID sets the "request_id" field.
func (u *AccessUpsertBulk) SetRequestID(v uuid.UUID) *AccessUpsertBulk {
	return u.Update(func(s *AccessUpsert) {
		s.SetRequestID(v)
	})
}

// UpdateRequestID sets the "request_id" field to the value that was provided on create.
func (u *AccessUpsertBulk) UpdateRequestID() *AccessUpsertBulk {
	return u.Update(func(s *AccessUpsert) {
		s.UpdateRequestID()
	})
}

// Exec executes the query.
func (u *AccessUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AccessCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AccessCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AccessUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
