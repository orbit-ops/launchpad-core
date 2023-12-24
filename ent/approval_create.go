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
	"github.com/orbit-ops/launchpad-core/ent/approval"
	"github.com/orbit-ops/launchpad-core/ent/request"
)

// ApprovalCreate is the builder for creating a Approval entity.
type ApprovalCreate struct {
	config
	mutation *ApprovalMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetPerson sets the "person" field.
func (ac *ApprovalCreate) SetPerson(s string) *ApprovalCreate {
	ac.mutation.SetPerson(s)
	return ac
}

// SetApprovedTime sets the "approved_time" field.
func (ac *ApprovalCreate) SetApprovedTime(t time.Time) *ApprovalCreate {
	ac.mutation.SetApprovedTime(t)
	return ac
}

// SetApproved sets the "approved" field.
func (ac *ApprovalCreate) SetApproved(b bool) *ApprovalCreate {
	ac.mutation.SetApproved(b)
	return ac
}

// SetRevoked sets the "revoked" field.
func (ac *ApprovalCreate) SetRevoked(b bool) *ApprovalCreate {
	ac.mutation.SetRevoked(b)
	return ac
}

// SetNillableRevoked sets the "revoked" field if the given value is not nil.
func (ac *ApprovalCreate) SetNillableRevoked(b *bool) *ApprovalCreate {
	if b != nil {
		ac.SetRevoked(*b)
	}
	return ac
}

// SetRevokedTime sets the "revoked_time" field.
func (ac *ApprovalCreate) SetRevokedTime(t time.Time) *ApprovalCreate {
	ac.mutation.SetRevokedTime(t)
	return ac
}

// SetNillableRevokedTime sets the "revoked_time" field if the given value is not nil.
func (ac *ApprovalCreate) SetNillableRevokedTime(t *time.Time) *ApprovalCreate {
	if t != nil {
		ac.SetRevokedTime(*t)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *ApprovalCreate) SetID(u uuid.UUID) *ApprovalCreate {
	ac.mutation.SetID(u)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *ApprovalCreate) SetNillableID(u *uuid.UUID) *ApprovalCreate {
	if u != nil {
		ac.SetID(*u)
	}
	return ac
}

// SetRequestID sets the "request" edge to the Request entity by ID.
func (ac *ApprovalCreate) SetRequestID(id uuid.UUID) *ApprovalCreate {
	ac.mutation.SetRequestID(id)
	return ac
}

// SetRequest sets the "request" edge to the Request entity.
func (ac *ApprovalCreate) SetRequest(r *Request) *ApprovalCreate {
	return ac.SetRequestID(r.ID)
}

// AddAccesIDs adds the "access" edge to the Access entity by IDs.
func (ac *ApprovalCreate) AddAccesIDs(ids ...uuid.UUID) *ApprovalCreate {
	ac.mutation.AddAccesIDs(ids...)
	return ac
}

// AddAccess adds the "access" edges to the Access entity.
func (ac *ApprovalCreate) AddAccess(a ...*Access) *ApprovalCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ac.AddAccesIDs(ids...)
}

// Mutation returns the ApprovalMutation object of the builder.
func (ac *ApprovalCreate) Mutation() *ApprovalMutation {
	return ac.mutation
}

// Save creates the Approval in the database.
func (ac *ApprovalCreate) Save(ctx context.Context) (*Approval, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *ApprovalCreate) SaveX(ctx context.Context) *Approval {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *ApprovalCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *ApprovalCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *ApprovalCreate) defaults() {
	if _, ok := ac.mutation.Revoked(); !ok {
		v := approval.DefaultRevoked
		ac.mutation.SetRevoked(v)
	}
	if _, ok := ac.mutation.ID(); !ok {
		v := approval.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *ApprovalCreate) check() error {
	if _, ok := ac.mutation.Person(); !ok {
		return &ValidationError{Name: "person", err: errors.New(`ent: missing required field "Approval.person"`)}
	}
	if _, ok := ac.mutation.ApprovedTime(); !ok {
		return &ValidationError{Name: "approved_time", err: errors.New(`ent: missing required field "Approval.approved_time"`)}
	}
	if _, ok := ac.mutation.Approved(); !ok {
		return &ValidationError{Name: "approved", err: errors.New(`ent: missing required field "Approval.approved"`)}
	}
	if _, ok := ac.mutation.Revoked(); !ok {
		return &ValidationError{Name: "revoked", err: errors.New(`ent: missing required field "Approval.revoked"`)}
	}
	if _, ok := ac.mutation.RequestID(); !ok {
		return &ValidationError{Name: "request", err: errors.New(`ent: missing required edge "Approval.request"`)}
	}
	return nil
}

func (ac *ApprovalCreate) sqlSave(ctx context.Context) (*Approval, error) {
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

func (ac *ApprovalCreate) createSpec() (*Approval, *sqlgraph.CreateSpec) {
	var (
		_node = &Approval{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(approval.Table, sqlgraph.NewFieldSpec(approval.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = ac.conflict
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ac.mutation.Person(); ok {
		_spec.SetField(approval.FieldPerson, field.TypeString, value)
		_node.Person = value
	}
	if value, ok := ac.mutation.ApprovedTime(); ok {
		_spec.SetField(approval.FieldApprovedTime, field.TypeTime, value)
		_node.ApprovedTime = value
	}
	if value, ok := ac.mutation.Approved(); ok {
		_spec.SetField(approval.FieldApproved, field.TypeBool, value)
		_node.Approved = value
	}
	if value, ok := ac.mutation.Revoked(); ok {
		_spec.SetField(approval.FieldRevoked, field.TypeBool, value)
		_node.Revoked = value
	}
	if value, ok := ac.mutation.RevokedTime(); ok {
		_spec.SetField(approval.FieldRevokedTime, field.TypeTime, value)
		_node.RevokedTime = value
	}
	if nodes := ac.mutation.RequestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   approval.RequestTable,
			Columns: []string{approval.RequestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(request.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.approval_request = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.AccessIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   approval.AccessTable,
			Columns: approval.AccessPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(access.FieldID, field.TypeUUID),
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
//	client.Approval.Create().
//		SetPerson(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ApprovalUpsert) {
//			SetPerson(v+v).
//		}).
//		Exec(ctx)
func (ac *ApprovalCreate) OnConflict(opts ...sql.ConflictOption) *ApprovalUpsertOne {
	ac.conflict = opts
	return &ApprovalUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Approval.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ac *ApprovalCreate) OnConflictColumns(columns ...string) *ApprovalUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &ApprovalUpsertOne{
		create: ac,
	}
}

type (
	// ApprovalUpsertOne is the builder for "upsert"-ing
	//  one Approval node.
	ApprovalUpsertOne struct {
		create *ApprovalCreate
	}

	// ApprovalUpsert is the "OnConflict" setter.
	ApprovalUpsert struct {
		*sql.UpdateSet
	}
)

// SetRevoked sets the "revoked" field.
func (u *ApprovalUpsert) SetRevoked(v bool) *ApprovalUpsert {
	u.Set(approval.FieldRevoked, v)
	return u
}

// UpdateRevoked sets the "revoked" field to the value that was provided on create.
func (u *ApprovalUpsert) UpdateRevoked() *ApprovalUpsert {
	u.SetExcluded(approval.FieldRevoked)
	return u
}

// SetRevokedTime sets the "revoked_time" field.
func (u *ApprovalUpsert) SetRevokedTime(v time.Time) *ApprovalUpsert {
	u.Set(approval.FieldRevokedTime, v)
	return u
}

// UpdateRevokedTime sets the "revoked_time" field to the value that was provided on create.
func (u *ApprovalUpsert) UpdateRevokedTime() *ApprovalUpsert {
	u.SetExcluded(approval.FieldRevokedTime)
	return u
}

// ClearRevokedTime clears the value of the "revoked_time" field.
func (u *ApprovalUpsert) ClearRevokedTime() *ApprovalUpsert {
	u.SetNull(approval.FieldRevokedTime)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Approval.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(approval.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ApprovalUpsertOne) UpdateNewValues() *ApprovalUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(approval.FieldID)
		}
		if _, exists := u.create.mutation.Person(); exists {
			s.SetIgnore(approval.FieldPerson)
		}
		if _, exists := u.create.mutation.ApprovedTime(); exists {
			s.SetIgnore(approval.FieldApprovedTime)
		}
		if _, exists := u.create.mutation.Approved(); exists {
			s.SetIgnore(approval.FieldApproved)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Approval.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ApprovalUpsertOne) Ignore() *ApprovalUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ApprovalUpsertOne) DoNothing() *ApprovalUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ApprovalCreate.OnConflict
// documentation for more info.
func (u *ApprovalUpsertOne) Update(set func(*ApprovalUpsert)) *ApprovalUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ApprovalUpsert{UpdateSet: update})
	}))
	return u
}

// SetRevoked sets the "revoked" field.
func (u *ApprovalUpsertOne) SetRevoked(v bool) *ApprovalUpsertOne {
	return u.Update(func(s *ApprovalUpsert) {
		s.SetRevoked(v)
	})
}

// UpdateRevoked sets the "revoked" field to the value that was provided on create.
func (u *ApprovalUpsertOne) UpdateRevoked() *ApprovalUpsertOne {
	return u.Update(func(s *ApprovalUpsert) {
		s.UpdateRevoked()
	})
}

// SetRevokedTime sets the "revoked_time" field.
func (u *ApprovalUpsertOne) SetRevokedTime(v time.Time) *ApprovalUpsertOne {
	return u.Update(func(s *ApprovalUpsert) {
		s.SetRevokedTime(v)
	})
}

// UpdateRevokedTime sets the "revoked_time" field to the value that was provided on create.
func (u *ApprovalUpsertOne) UpdateRevokedTime() *ApprovalUpsertOne {
	return u.Update(func(s *ApprovalUpsert) {
		s.UpdateRevokedTime()
	})
}

// ClearRevokedTime clears the value of the "revoked_time" field.
func (u *ApprovalUpsertOne) ClearRevokedTime() *ApprovalUpsertOne {
	return u.Update(func(s *ApprovalUpsert) {
		s.ClearRevokedTime()
	})
}

// Exec executes the query.
func (u *ApprovalUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ApprovalCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ApprovalUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ApprovalUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ApprovalUpsertOne.ID is not supported by MySQL driver. Use ApprovalUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ApprovalUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ApprovalCreateBulk is the builder for creating many Approval entities in bulk.
type ApprovalCreateBulk struct {
	config
	err      error
	builders []*ApprovalCreate
	conflict []sql.ConflictOption
}

// Save creates the Approval entities in the database.
func (acb *ApprovalCreateBulk) Save(ctx context.Context) ([]*Approval, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Approval, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ApprovalMutation)
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
func (acb *ApprovalCreateBulk) SaveX(ctx context.Context) []*Approval {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *ApprovalCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *ApprovalCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Approval.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ApprovalUpsert) {
//			SetPerson(v+v).
//		}).
//		Exec(ctx)
func (acb *ApprovalCreateBulk) OnConflict(opts ...sql.ConflictOption) *ApprovalUpsertBulk {
	acb.conflict = opts
	return &ApprovalUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Approval.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acb *ApprovalCreateBulk) OnConflictColumns(columns ...string) *ApprovalUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &ApprovalUpsertBulk{
		create: acb,
	}
}

// ApprovalUpsertBulk is the builder for "upsert"-ing
// a bulk of Approval nodes.
type ApprovalUpsertBulk struct {
	create *ApprovalCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Approval.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(approval.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ApprovalUpsertBulk) UpdateNewValues() *ApprovalUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(approval.FieldID)
			}
			if _, exists := b.mutation.Person(); exists {
				s.SetIgnore(approval.FieldPerson)
			}
			if _, exists := b.mutation.ApprovedTime(); exists {
				s.SetIgnore(approval.FieldApprovedTime)
			}
			if _, exists := b.mutation.Approved(); exists {
				s.SetIgnore(approval.FieldApproved)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Approval.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ApprovalUpsertBulk) Ignore() *ApprovalUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ApprovalUpsertBulk) DoNothing() *ApprovalUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ApprovalCreateBulk.OnConflict
// documentation for more info.
func (u *ApprovalUpsertBulk) Update(set func(*ApprovalUpsert)) *ApprovalUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ApprovalUpsert{UpdateSet: update})
	}))
	return u
}

// SetRevoked sets the "revoked" field.
func (u *ApprovalUpsertBulk) SetRevoked(v bool) *ApprovalUpsertBulk {
	return u.Update(func(s *ApprovalUpsert) {
		s.SetRevoked(v)
	})
}

// UpdateRevoked sets the "revoked" field to the value that was provided on create.
func (u *ApprovalUpsertBulk) UpdateRevoked() *ApprovalUpsertBulk {
	return u.Update(func(s *ApprovalUpsert) {
		s.UpdateRevoked()
	})
}

// SetRevokedTime sets the "revoked_time" field.
func (u *ApprovalUpsertBulk) SetRevokedTime(v time.Time) *ApprovalUpsertBulk {
	return u.Update(func(s *ApprovalUpsert) {
		s.SetRevokedTime(v)
	})
}

// UpdateRevokedTime sets the "revoked_time" field to the value that was provided on create.
func (u *ApprovalUpsertBulk) UpdateRevokedTime() *ApprovalUpsertBulk {
	return u.Update(func(s *ApprovalUpsert) {
		s.UpdateRevokedTime()
	})
}

// ClearRevokedTime clears the value of the "revoked_time" field.
func (u *ApprovalUpsertBulk) ClearRevokedTime() *ApprovalUpsertBulk {
	return u.Update(func(s *ApprovalUpsert) {
		s.ClearRevokedTime()
	})
}

// Exec executes the query.
func (u *ApprovalUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ApprovalCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ApprovalCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ApprovalUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
