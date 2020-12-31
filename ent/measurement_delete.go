// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/vmkevv/suprat-api/ent/measurement"
	"github.com/vmkevv/suprat-api/ent/predicate"
)

// MeasurementDelete is the builder for deleting a Measurement entity.
type MeasurementDelete struct {
	config
	hooks    []Hook
	mutation *MeasurementMutation
}

// Where adds a new predicate to the delete builder.
func (md *MeasurementDelete) Where(ps ...predicate.Measurement) *MeasurementDelete {
	md.mutation.predicates = append(md.mutation.predicates, ps...)
	return md
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (md *MeasurementDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(md.hooks) == 0 {
		affected, err = md.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MeasurementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			md.mutation = mutation
			affected, err = md.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(md.hooks) - 1; i >= 0; i-- {
			mut = md.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, md.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (md *MeasurementDelete) ExecX(ctx context.Context) int {
	n, err := md.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (md *MeasurementDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: measurement.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: measurement.FieldID,
			},
		},
	}
	if ps := md.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, md.driver, _spec)
}

// MeasurementDeleteOne is the builder for deleting a single Measurement entity.
type MeasurementDeleteOne struct {
	md *MeasurementDelete
}

// Exec executes the deletion query.
func (mdo *MeasurementDeleteOne) Exec(ctx context.Context) error {
	n, err := mdo.md.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{measurement.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mdo *MeasurementDeleteOne) ExecX(ctx context.Context) {
	mdo.md.ExecX(ctx)
}
