// Code generated by entc, DO NOT EDIT.

package record

import (
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/vmkevv/suprat-api/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Start applies equality check predicate on the "start" field. It's identical to StartEQ.
func Start(v time.Time) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStart), v))
	})
}

// End applies equality check predicate on the "end" field. It's identical to EndEQ.
func End(v time.Time) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnd), v))
	})
}

// StartEQ applies the EQ predicate on the "start" field.
func StartEQ(v time.Time) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStart), v))
	})
}

// StartNEQ applies the NEQ predicate on the "start" field.
func StartNEQ(v time.Time) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStart), v))
	})
}

// StartIn applies the In predicate on the "start" field.
func StartIn(vs ...time.Time) predicate.Record {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Record(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStart), v...))
	})
}

// StartNotIn applies the NotIn predicate on the "start" field.
func StartNotIn(vs ...time.Time) predicate.Record {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Record(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStart), v...))
	})
}

// StartGT applies the GT predicate on the "start" field.
func StartGT(v time.Time) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStart), v))
	})
}

// StartGTE applies the GTE predicate on the "start" field.
func StartGTE(v time.Time) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStart), v))
	})
}

// StartLT applies the LT predicate on the "start" field.
func StartLT(v time.Time) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStart), v))
	})
}

// StartLTE applies the LTE predicate on the "start" field.
func StartLTE(v time.Time) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStart), v))
	})
}

// EndEQ applies the EQ predicate on the "end" field.
func EndEQ(v time.Time) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnd), v))
	})
}

// EndNEQ applies the NEQ predicate on the "end" field.
func EndNEQ(v time.Time) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEnd), v))
	})
}

// EndIn applies the In predicate on the "end" field.
func EndIn(vs ...time.Time) predicate.Record {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Record(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEnd), v...))
	})
}

// EndNotIn applies the NotIn predicate on the "end" field.
func EndNotIn(vs ...time.Time) predicate.Record {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Record(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEnd), v...))
	})
}

// EndGT applies the GT predicate on the "end" field.
func EndGT(v time.Time) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEnd), v))
	})
}

// EndGTE applies the GTE predicate on the "end" field.
func EndGTE(v time.Time) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEnd), v))
	})
}

// EndLT applies the LT predicate on the "end" field.
func EndLT(v time.Time) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEnd), v))
	})
}

// EndLTE applies the LTE predicate on the "end" field.
func EndLTE(v time.Time) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEnd), v))
	})
}

// HasActivity applies the HasEdge predicate on the "activity" edge.
func HasActivity() predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ActivityTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ActivityTable, ActivityColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasActivityWith applies the HasEdge predicate on the "activity" edge with a given conditions (other predicates).
func HasActivityWith(preds ...predicate.Activity) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ActivityInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ActivityTable, ActivityColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMeasures applies the HasEdge predicate on the "measures" edge.
func HasMeasures() predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MeasuresTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MeasuresTable, MeasuresColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMeasuresWith applies the HasEdge predicate on the "measures" edge with a given conditions (other predicates).
func HasMeasuresWith(preds ...predicate.Measure) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MeasuresInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MeasuresTable, MeasuresColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Record) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Record) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Record) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		p(s.Not())
	})
}
