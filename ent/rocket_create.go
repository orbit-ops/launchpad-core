// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/orbit-ops/launchpad-core/ent/mission"
	"github.com/orbit-ops/launchpad-core/ent/rocket"
)

// RocketCreate is the builder for creating a Rocket entity.
type RocketCreate struct {
	config
	mutation *RocketMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetDescription sets the "description" field.
func (rc *RocketCreate) SetDescription(s string) *RocketCreate {
	rc.mutation.SetDescription(s)
	return rc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (rc *RocketCreate) SetNillableDescription(s *string) *RocketCreate {
	if s != nil {
		rc.SetDescription(*s)
	}
	return rc
}

// SetImage sets the "image" field.
func (rc *RocketCreate) SetImage(s string) *RocketCreate {
	rc.mutation.SetImage(s)
	return rc
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (rc *RocketCreate) SetNillableImage(s *string) *RocketCreate {
	if s != nil {
		rc.SetImage(*s)
	}
	return rc
}

// SetZip sets the "zip" field.
func (rc *RocketCreate) SetZip(s string) *RocketCreate {
	rc.mutation.SetZip(s)
	return rc
}

// SetNillableZip sets the "zip" field if the given value is not nil.
func (rc *RocketCreate) SetNillableZip(s *string) *RocketCreate {
	if s != nil {
		rc.SetZip(*s)
	}
	return rc
}

// SetConfig sets the "config" field.
func (rc *RocketCreate) SetConfig(m map[string]string) *RocketCreate {
	rc.mutation.SetConfig(m)
	return rc
}

// SetID sets the "id" field.
func (rc *RocketCreate) SetID(s string) *RocketCreate {
	rc.mutation.SetID(s)
	return rc
}

// AddMissionIDs adds the "missions" edge to the Mission entity by IDs.
func (rc *RocketCreate) AddMissionIDs(ids ...string) *RocketCreate {
	rc.mutation.AddMissionIDs(ids...)
	return rc
}

// AddMissions adds the "missions" edges to the Mission entity.
func (rc *RocketCreate) AddMissions(m ...*Mission) *RocketCreate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return rc.AddMissionIDs(ids...)
}

// Mutation returns the RocketMutation object of the builder.
func (rc *RocketCreate) Mutation() *RocketMutation {
	return rc.mutation
}

// Save creates the Rocket in the database.
func (rc *RocketCreate) Save(ctx context.Context) (*Rocket, error) {
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RocketCreate) SaveX(ctx context.Context) *Rocket {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RocketCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RocketCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RocketCreate) check() error {
	if v, ok := rc.mutation.Image(); ok {
		if err := rocket.ImageValidator(v); err != nil {
			return &ValidationError{Name: "image", err: fmt.Errorf(`ent: validator failed for field "Rocket.image": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Config(); !ok {
		return &ValidationError{Name: "config", err: errors.New(`ent: missing required field "Rocket.config"`)}
	}
	return nil
}

func (rc *RocketCreate) sqlSave(ctx context.Context) (*Rocket, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Rocket.ID type: %T", _spec.ID.Value)
		}
	}
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RocketCreate) createSpec() (*Rocket, *sqlgraph.CreateSpec) {
	var (
		_node = &Rocket{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(rocket.Table, sqlgraph.NewFieldSpec(rocket.FieldID, field.TypeString))
	)
	_spec.OnConflict = rc.conflict
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.Description(); ok {
		_spec.SetField(rocket.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := rc.mutation.Image(); ok {
		_spec.SetField(rocket.FieldImage, field.TypeString, value)
		_node.Image = value
	}
	if value, ok := rc.mutation.Zip(); ok {
		_spec.SetField(rocket.FieldZip, field.TypeString, value)
		_node.Zip = value
	}
	if value, ok := rc.mutation.Config(); ok {
		_spec.SetField(rocket.FieldConfig, field.TypeJSON, value)
		_node.Config = value
	}
	if nodes := rc.mutation.MissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   rocket.MissionsTable,
			Columns: rocket.MissionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(mission.FieldID, field.TypeString),
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
//	client.Rocket.Create().
//		SetDescription(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RocketUpsert) {
//			SetDescription(v+v).
//		}).
//		Exec(ctx)
func (rc *RocketCreate) OnConflict(opts ...sql.ConflictOption) *RocketUpsertOne {
	rc.conflict = opts
	return &RocketUpsertOne{
		create: rc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Rocket.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rc *RocketCreate) OnConflictColumns(columns ...string) *RocketUpsertOne {
	rc.conflict = append(rc.conflict, sql.ConflictColumns(columns...))
	return &RocketUpsertOne{
		create: rc,
	}
}

type (
	// RocketUpsertOne is the builder for "upsert"-ing
	//  one Rocket node.
	RocketUpsertOne struct {
		create *RocketCreate
	}

	// RocketUpsert is the "OnConflict" setter.
	RocketUpsert struct {
		*sql.UpdateSet
	}
)

// SetDescription sets the "description" field.
func (u *RocketUpsert) SetDescription(v string) *RocketUpsert {
	u.Set(rocket.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *RocketUpsert) UpdateDescription() *RocketUpsert {
	u.SetExcluded(rocket.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *RocketUpsert) ClearDescription() *RocketUpsert {
	u.SetNull(rocket.FieldDescription)
	return u
}

// SetImage sets the "image" field.
func (u *RocketUpsert) SetImage(v string) *RocketUpsert {
	u.Set(rocket.FieldImage, v)
	return u
}

// UpdateImage sets the "image" field to the value that was provided on create.
func (u *RocketUpsert) UpdateImage() *RocketUpsert {
	u.SetExcluded(rocket.FieldImage)
	return u
}

// ClearImage clears the value of the "image" field.
func (u *RocketUpsert) ClearImage() *RocketUpsert {
	u.SetNull(rocket.FieldImage)
	return u
}

// SetZip sets the "zip" field.
func (u *RocketUpsert) SetZip(v string) *RocketUpsert {
	u.Set(rocket.FieldZip, v)
	return u
}

// UpdateZip sets the "zip" field to the value that was provided on create.
func (u *RocketUpsert) UpdateZip() *RocketUpsert {
	u.SetExcluded(rocket.FieldZip)
	return u
}

// ClearZip clears the value of the "zip" field.
func (u *RocketUpsert) ClearZip() *RocketUpsert {
	u.SetNull(rocket.FieldZip)
	return u
}

// SetConfig sets the "config" field.
func (u *RocketUpsert) SetConfig(v map[string]string) *RocketUpsert {
	u.Set(rocket.FieldConfig, v)
	return u
}

// UpdateConfig sets the "config" field to the value that was provided on create.
func (u *RocketUpsert) UpdateConfig() *RocketUpsert {
	u.SetExcluded(rocket.FieldConfig)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Rocket.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(rocket.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *RocketUpsertOne) UpdateNewValues() *RocketUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(rocket.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Rocket.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *RocketUpsertOne) Ignore() *RocketUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RocketUpsertOne) DoNothing() *RocketUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RocketCreate.OnConflict
// documentation for more info.
func (u *RocketUpsertOne) Update(set func(*RocketUpsert)) *RocketUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RocketUpsert{UpdateSet: update})
	}))
	return u
}

// SetDescription sets the "description" field.
func (u *RocketUpsertOne) SetDescription(v string) *RocketUpsertOne {
	return u.Update(func(s *RocketUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *RocketUpsertOne) UpdateDescription() *RocketUpsertOne {
	return u.Update(func(s *RocketUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *RocketUpsertOne) ClearDescription() *RocketUpsertOne {
	return u.Update(func(s *RocketUpsert) {
		s.ClearDescription()
	})
}

// SetImage sets the "image" field.
func (u *RocketUpsertOne) SetImage(v string) *RocketUpsertOne {
	return u.Update(func(s *RocketUpsert) {
		s.SetImage(v)
	})
}

// UpdateImage sets the "image" field to the value that was provided on create.
func (u *RocketUpsertOne) UpdateImage() *RocketUpsertOne {
	return u.Update(func(s *RocketUpsert) {
		s.UpdateImage()
	})
}

// ClearImage clears the value of the "image" field.
func (u *RocketUpsertOne) ClearImage() *RocketUpsertOne {
	return u.Update(func(s *RocketUpsert) {
		s.ClearImage()
	})
}

// SetZip sets the "zip" field.
func (u *RocketUpsertOne) SetZip(v string) *RocketUpsertOne {
	return u.Update(func(s *RocketUpsert) {
		s.SetZip(v)
	})
}

// UpdateZip sets the "zip" field to the value that was provided on create.
func (u *RocketUpsertOne) UpdateZip() *RocketUpsertOne {
	return u.Update(func(s *RocketUpsert) {
		s.UpdateZip()
	})
}

// ClearZip clears the value of the "zip" field.
func (u *RocketUpsertOne) ClearZip() *RocketUpsertOne {
	return u.Update(func(s *RocketUpsert) {
		s.ClearZip()
	})
}

// SetConfig sets the "config" field.
func (u *RocketUpsertOne) SetConfig(v map[string]string) *RocketUpsertOne {
	return u.Update(func(s *RocketUpsert) {
		s.SetConfig(v)
	})
}

// UpdateConfig sets the "config" field to the value that was provided on create.
func (u *RocketUpsertOne) UpdateConfig() *RocketUpsertOne {
	return u.Update(func(s *RocketUpsert) {
		s.UpdateConfig()
	})
}

// Exec executes the query.
func (u *RocketUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RocketCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RocketUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *RocketUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: RocketUpsertOne.ID is not supported by MySQL driver. Use RocketUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *RocketUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// RocketCreateBulk is the builder for creating many Rocket entities in bulk.
type RocketCreateBulk struct {
	config
	err      error
	builders []*RocketCreate
	conflict []sql.ConflictOption
}

// Save creates the Rocket entities in the database.
func (rcb *RocketCreateBulk) Save(ctx context.Context) ([]*Rocket, error) {
	if rcb.err != nil {
		return nil, rcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Rocket, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RocketMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = rcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RocketCreateBulk) SaveX(ctx context.Context) []*Rocket {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RocketCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RocketCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Rocket.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RocketUpsert) {
//			SetDescription(v+v).
//		}).
//		Exec(ctx)
func (rcb *RocketCreateBulk) OnConflict(opts ...sql.ConflictOption) *RocketUpsertBulk {
	rcb.conflict = opts
	return &RocketUpsertBulk{
		create: rcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Rocket.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rcb *RocketCreateBulk) OnConflictColumns(columns ...string) *RocketUpsertBulk {
	rcb.conflict = append(rcb.conflict, sql.ConflictColumns(columns...))
	return &RocketUpsertBulk{
		create: rcb,
	}
}

// RocketUpsertBulk is the builder for "upsert"-ing
// a bulk of Rocket nodes.
type RocketUpsertBulk struct {
	create *RocketCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Rocket.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(rocket.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *RocketUpsertBulk) UpdateNewValues() *RocketUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(rocket.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Rocket.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *RocketUpsertBulk) Ignore() *RocketUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RocketUpsertBulk) DoNothing() *RocketUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RocketCreateBulk.OnConflict
// documentation for more info.
func (u *RocketUpsertBulk) Update(set func(*RocketUpsert)) *RocketUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RocketUpsert{UpdateSet: update})
	}))
	return u
}

// SetDescription sets the "description" field.
func (u *RocketUpsertBulk) SetDescription(v string) *RocketUpsertBulk {
	return u.Update(func(s *RocketUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *RocketUpsertBulk) UpdateDescription() *RocketUpsertBulk {
	return u.Update(func(s *RocketUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *RocketUpsertBulk) ClearDescription() *RocketUpsertBulk {
	return u.Update(func(s *RocketUpsert) {
		s.ClearDescription()
	})
}

// SetImage sets the "image" field.
func (u *RocketUpsertBulk) SetImage(v string) *RocketUpsertBulk {
	return u.Update(func(s *RocketUpsert) {
		s.SetImage(v)
	})
}

// UpdateImage sets the "image" field to the value that was provided on create.
func (u *RocketUpsertBulk) UpdateImage() *RocketUpsertBulk {
	return u.Update(func(s *RocketUpsert) {
		s.UpdateImage()
	})
}

// ClearImage clears the value of the "image" field.
func (u *RocketUpsertBulk) ClearImage() *RocketUpsertBulk {
	return u.Update(func(s *RocketUpsert) {
		s.ClearImage()
	})
}

// SetZip sets the "zip" field.
func (u *RocketUpsertBulk) SetZip(v string) *RocketUpsertBulk {
	return u.Update(func(s *RocketUpsert) {
		s.SetZip(v)
	})
}

// UpdateZip sets the "zip" field to the value that was provided on create.
func (u *RocketUpsertBulk) UpdateZip() *RocketUpsertBulk {
	return u.Update(func(s *RocketUpsert) {
		s.UpdateZip()
	})
}

// ClearZip clears the value of the "zip" field.
func (u *RocketUpsertBulk) ClearZip() *RocketUpsertBulk {
	return u.Update(func(s *RocketUpsert) {
		s.ClearZip()
	})
}

// SetConfig sets the "config" field.
func (u *RocketUpsertBulk) SetConfig(v map[string]string) *RocketUpsertBulk {
	return u.Update(func(s *RocketUpsert) {
		s.SetConfig(v)
	})
}

// UpdateConfig sets the "config" field to the value that was provided on create.
func (u *RocketUpsertBulk) UpdateConfig() *RocketUpsertBulk {
	return u.Update(func(s *RocketUpsert) {
		s.UpdateConfig()
	})
}

// Exec executes the query.
func (u *RocketUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the RocketCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RocketCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RocketUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
