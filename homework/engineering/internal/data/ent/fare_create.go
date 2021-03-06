// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/webmin7761/go-school/homework/engineering/internal/data/ent/fare"
)

// FareCreate is the builder for creating a Fare entity.
type FareCreate struct {
	config
	mutation *FareMutation
	hooks    []Hook
}

// SetOrgAirport sets the "org_airport" field.
func (fc *FareCreate) SetOrgAirport(s string) *FareCreate {
	fc.mutation.SetOrgAirport(s)
	return fc
}

// SetArrAirport sets the "arr_airport" field.
func (fc *FareCreate) SetArrAirport(s string) *FareCreate {
	fc.mutation.SetArrAirport(s)
	return fc
}

// SetPassageType sets the "passage_type" field.
func (fc *FareCreate) SetPassageType(s string) *FareCreate {
	fc.mutation.SetPassageType(s)
	return fc
}

// SetFirstTravelDate sets the "first_travel_date" field.
func (fc *FareCreate) SetFirstTravelDate(t time.Time) *FareCreate {
	fc.mutation.SetFirstTravelDate(t)
	return fc
}

// SetNillableFirstTravelDate sets the "first_travel_date" field if the given value is not nil.
func (fc *FareCreate) SetNillableFirstTravelDate(t *time.Time) *FareCreate {
	if t != nil {
		fc.SetFirstTravelDate(*t)
	}
	return fc
}

// SetLastTravelDate sets the "last_travel_date" field.
func (fc *FareCreate) SetLastTravelDate(t time.Time) *FareCreate {
	fc.mutation.SetLastTravelDate(t)
	return fc
}

// SetNillableLastTravelDate sets the "last_travel_date" field if the given value is not nil.
func (fc *FareCreate) SetNillableLastTravelDate(t *time.Time) *FareCreate {
	if t != nil {
		fc.SetLastTravelDate(*t)
	}
	return fc
}

// SetAmount sets the "amount" field.
func (fc *FareCreate) SetAmount(f float64) *FareCreate {
	fc.mutation.SetAmount(f)
	return fc
}

// SetCreatedAt sets the "created_at" field.
func (fc *FareCreate) SetCreatedAt(t time.Time) *FareCreate {
	fc.mutation.SetCreatedAt(t)
	return fc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fc *FareCreate) SetNillableCreatedAt(t *time.Time) *FareCreate {
	if t != nil {
		fc.SetCreatedAt(*t)
	}
	return fc
}

// SetUpdatedAt sets the "updated_at" field.
func (fc *FareCreate) SetUpdatedAt(t time.Time) *FareCreate {
	fc.mutation.SetUpdatedAt(t)
	return fc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fc *FareCreate) SetNillableUpdatedAt(t *time.Time) *FareCreate {
	if t != nil {
		fc.SetUpdatedAt(*t)
	}
	return fc
}

// SetID sets the "id" field.
func (fc *FareCreate) SetID(i int64) *FareCreate {
	fc.mutation.SetID(i)
	return fc
}

// Mutation returns the FareMutation object of the builder.
func (fc *FareCreate) Mutation() *FareMutation {
	return fc.mutation
}

// Save creates the Fare in the database.
func (fc *FareCreate) Save(ctx context.Context) (*Fare, error) {
	var (
		err  error
		node *Fare
	)
	fc.defaults()
	if len(fc.hooks) == 0 {
		if err = fc.check(); err != nil {
			return nil, err
		}
		node, err = fc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FareMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fc.check(); err != nil {
				return nil, err
			}
			fc.mutation = mutation
			node, err = fc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fc.hooks) - 1; i >= 0; i-- {
			mut = fc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FareCreate) SaveX(ctx context.Context) *Fare {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (fc *FareCreate) defaults() {
	if _, ok := fc.mutation.FirstTravelDate(); !ok {
		v := fare.DefaultFirstTravelDate()
		fc.mutation.SetFirstTravelDate(v)
	}
	if _, ok := fc.mutation.LastTravelDate(); !ok {
		v := fare.DefaultLastTravelDate()
		fc.mutation.SetLastTravelDate(v)
	}
	if _, ok := fc.mutation.CreatedAt(); !ok {
		v := fare.DefaultCreatedAt()
		fc.mutation.SetCreatedAt(v)
	}
	if _, ok := fc.mutation.UpdatedAt(); !ok {
		v := fare.DefaultUpdatedAt()
		fc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FareCreate) check() error {
	if _, ok := fc.mutation.OrgAirport(); !ok {
		return &ValidationError{Name: "org_airport", err: errors.New("ent: missing required field \"org_airport\"")}
	}
	if _, ok := fc.mutation.ArrAirport(); !ok {
		return &ValidationError{Name: "arr_airport", err: errors.New("ent: missing required field \"arr_airport\"")}
	}
	if _, ok := fc.mutation.PassageType(); !ok {
		return &ValidationError{Name: "passage_type", err: errors.New("ent: missing required field \"passage_type\"")}
	}
	if _, ok := fc.mutation.FirstTravelDate(); !ok {
		return &ValidationError{Name: "first_travel_date", err: errors.New("ent: missing required field \"first_travel_date\"")}
	}
	if _, ok := fc.mutation.LastTravelDate(); !ok {
		return &ValidationError{Name: "last_travel_date", err: errors.New("ent: missing required field \"last_travel_date\"")}
	}
	if _, ok := fc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New("ent: missing required field \"amount\"")}
	}
	if _, ok := fc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := fc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	return nil
}

func (fc *FareCreate) sqlSave(ctx context.Context) (*Fare, error) {
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	if _node.ID == 0 {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (fc *FareCreate) createSpec() (*Fare, *sqlgraph.CreateSpec) {
	var (
		_node = &Fare{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: fare.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: fare.FieldID,
			},
		}
	)
	if id, ok := fc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := fc.mutation.OrgAirport(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fare.FieldOrgAirport,
		})
		_node.OrgAirport = value
	}
	if value, ok := fc.mutation.ArrAirport(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fare.FieldArrAirport,
		})
		_node.ArrAirport = value
	}
	if value, ok := fc.mutation.PassageType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fare.FieldPassageType,
		})
		_node.PassageType = value
	}
	if value, ok := fc.mutation.FirstTravelDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: fare.FieldFirstTravelDate,
		})
		_node.FirstTravelDate = value
	}
	if value, ok := fc.mutation.LastTravelDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: fare.FieldLastTravelDate,
		})
		_node.LastTravelDate = value
	}
	if value, ok := fc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: fare.FieldAmount,
		})
		_node.Amount = value
	}
	if value, ok := fc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: fare.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := fc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: fare.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// FareCreateBulk is the builder for creating many Fare entities in bulk.
type FareCreateBulk struct {
	config
	builders []*FareCreate
}

// Save creates the Fare entities in the database.
func (fcb *FareCreateBulk) Save(ctx context.Context) ([]*Fare, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Fare, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FareMutation)
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
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				if nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
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
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FareCreateBulk) SaveX(ctx context.Context) []*Fare {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
