// Code generated by entc, DO NOT EDIT.

package basket

import (
	"ShoppingBasket/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Basket {
	return predicate.Basket(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Basket {
	return predicate.Basket(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Basket {
	return predicate.Basket(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Basket {
	return predicate.Basket(func(s *sql.Selector) {
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
func IDNotIn(ids ...uuid.UUID) predicate.Basket {
	return predicate.Basket(func(s *sql.Selector) {
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
func IDGT(id uuid.UUID) predicate.Basket {
	return predicate.Basket(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Basket {
	return predicate.Basket(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Basket {
	return predicate.Basket(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Basket {
	return predicate.Basket(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UserId applies equality check predicate on the "UserId" field. It's identical to UserIdEQ.
func UserId(v uuid.UUID) predicate.Basket {
	return predicate.Basket(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserId), v))
	})
}

// UserIdEQ applies the EQ predicate on the "UserId" field.
func UserIdEQ(v uuid.UUID) predicate.Basket {
	return predicate.Basket(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserId), v))
	})
}

// UserIdNEQ applies the NEQ predicate on the "UserId" field.
func UserIdNEQ(v uuid.UUID) predicate.Basket {
	return predicate.Basket(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserId), v))
	})
}

// UserIdIn applies the In predicate on the "UserId" field.
func UserIdIn(vs ...uuid.UUID) predicate.Basket {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Basket(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserId), v...))
	})
}

// UserIdNotIn applies the NotIn predicate on the "UserId" field.
func UserIdNotIn(vs ...uuid.UUID) predicate.Basket {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Basket(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserId), v...))
	})
}

// UserIdGT applies the GT predicate on the "UserId" field.
func UserIdGT(v uuid.UUID) predicate.Basket {
	return predicate.Basket(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserId), v))
	})
}

// UserIdGTE applies the GTE predicate on the "UserId" field.
func UserIdGTE(v uuid.UUID) predicate.Basket {
	return predicate.Basket(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserId), v))
	})
}

// UserIdLT applies the LT predicate on the "UserId" field.
func UserIdLT(v uuid.UUID) predicate.Basket {
	return predicate.Basket(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserId), v))
	})
}

// UserIdLTE applies the LTE predicate on the "UserId" field.
func UserIdLTE(v uuid.UUID) predicate.Basket {
	return predicate.Basket(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserId), v))
	})
}

// HasBasketLine applies the HasEdge predicate on the "BasketLine" edge.
func HasBasketLine() predicate.Basket {
	return predicate.Basket(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BasketLineTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BasketLineTable, BasketLineColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBasketLineWith applies the HasEdge predicate on the "BasketLine" edge with a given conditions (other predicates).
func HasBasketLineWith(preds ...predicate.BasketLine) predicate.Basket {
	return predicate.Basket(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BasketLineInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BasketLineTable, BasketLineColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Basket) predicate.Basket {
	return predicate.Basket(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Basket) predicate.Basket {
	return predicate.Basket(func(s *sql.Selector) {
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
func Not(p predicate.Basket) predicate.Basket {
	return predicate.Basket(func(s *sql.Selector) {
		p(s.Not())
	})
}
