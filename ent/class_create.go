// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-ent-graphql/ent/class"
	"go-ent-graphql/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ClassCreate is the builder for creating a Class entity.
type ClassCreate struct {
	config
	mutation *ClassMutation
	hooks    []Hook
}

// SetClassID sets the "class_id" field.
func (cc *ClassCreate) SetClassID(u uuid.UUID) *ClassCreate {
	cc.mutation.SetClassID(u)
	return cc
}

// SetName sets the "name" field.
func (cc *ClassCreate) SetName(s string) *ClassCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cc *ClassCreate) SetNillableName(s *string) *ClassCreate {
	if s != nil {
		cc.SetName(*s)
	}
	return cc
}

// SetCreatedAt sets the "created_at" field.
func (cc *ClassCreate) SetCreatedAt(t time.Time) *ClassCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *ClassCreate) SetNillableCreatedAt(t *time.Time) *ClassCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *ClassCreate) SetUpdatedAt(t time.Time) *ClassCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *ClassCreate) SetNillableUpdatedAt(t *time.Time) *ClassCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (cc *ClassCreate) AddUserIDs(ids ...int) *ClassCreate {
	cc.mutation.AddUserIDs(ids...)
	return cc
}

// AddUsers adds the "users" edges to the User entity.
func (cc *ClassCreate) AddUsers(u ...*User) *ClassCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cc.AddUserIDs(ids...)
}

// Mutation returns the ClassMutation object of the builder.
func (cc *ClassCreate) Mutation() *ClassMutation {
	return cc.mutation
}

// Save creates the Class in the database.
func (cc *ClassCreate) Save(ctx context.Context) (*Class, error) {
	var (
		err  error
		node *Class
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClassMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ClassCreate) SaveX(ctx context.Context) *Class {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ClassCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ClassCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ClassCreate) defaults() {
	if _, ok := cc.mutation.ClassID(); !ok {
		v := class.DefaultClassID()
		cc.mutation.SetClassID(v)
	}
	if _, ok := cc.mutation.Name(); !ok {
		v := class.DefaultName
		cc.mutation.SetName(v)
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := class.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ClassCreate) check() error {
	if _, ok := cc.mutation.ClassID(); !ok {
		return &ValidationError{Name: "class_id", err: errors.New(`ent: missing required field "class_id"`)}
	}
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	return nil
}

func (cc *ClassCreate) sqlSave(ctx context.Context) (*Class, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cc *ClassCreate) createSpec() (*Class, *sqlgraph.CreateSpec) {
	var (
		_node = &Class{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: class.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: class.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.ClassID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: class.FieldClassID,
		})
		_node.ClassID = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: class.FieldName,
		})
		_node.Name = value
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: class.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: class.FieldUpdatedAt,
		})
		_node.UpdatedAt = &value
	}
	if nodes := cc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   class.UsersTable,
			Columns: class.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ClassCreateBulk is the builder for creating many Class entities in bulk.
type ClassCreateBulk struct {
	config
	builders []*ClassCreate
}

// Save creates the Class entities in the database.
func (ccb *ClassCreateBulk) Save(ctx context.Context) ([]*Class, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Class, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ClassMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ClassCreateBulk) SaveX(ctx context.Context) []*Class {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ClassCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ClassCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
