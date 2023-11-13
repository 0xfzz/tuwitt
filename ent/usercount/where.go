// Code generated by ent, DO NOT EDIT.

package usercount

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/0xfzz/tuwitt/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UserCount {
	return predicate.UserCount(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UserCount {
	return predicate.UserCount(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UserCount {
	return predicate.UserCount(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UserCount {
	return predicate.UserCount(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.UserCount {
	return predicate.UserCount(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.UserCount {
	return predicate.UserCount(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UserCount {
	return predicate.UserCount(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UserCount {
	return predicate.UserCount(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UserCount {
	return predicate.UserCount(sql.FieldLTE(FieldID, id))
}

// FollowerCount applies equality check predicate on the "follower_count" field. It's identical to FollowerCountEQ.
func FollowerCount(v int) predicate.UserCount {
	return predicate.UserCount(sql.FieldEQ(FieldFollowerCount, v))
}

// FollowingsCount applies equality check predicate on the "followings_count" field. It's identical to FollowingsCountEQ.
func FollowingsCount(v int) predicate.UserCount {
	return predicate.UserCount(sql.FieldEQ(FieldFollowingsCount, v))
}

// FollowerCountEQ applies the EQ predicate on the "follower_count" field.
func FollowerCountEQ(v int) predicate.UserCount {
	return predicate.UserCount(sql.FieldEQ(FieldFollowerCount, v))
}

// FollowerCountNEQ applies the NEQ predicate on the "follower_count" field.
func FollowerCountNEQ(v int) predicate.UserCount {
	return predicate.UserCount(sql.FieldNEQ(FieldFollowerCount, v))
}

// FollowerCountIn applies the In predicate on the "follower_count" field.
func FollowerCountIn(vs ...int) predicate.UserCount {
	return predicate.UserCount(sql.FieldIn(FieldFollowerCount, vs...))
}

// FollowerCountNotIn applies the NotIn predicate on the "follower_count" field.
func FollowerCountNotIn(vs ...int) predicate.UserCount {
	return predicate.UserCount(sql.FieldNotIn(FieldFollowerCount, vs...))
}

// FollowerCountGT applies the GT predicate on the "follower_count" field.
func FollowerCountGT(v int) predicate.UserCount {
	return predicate.UserCount(sql.FieldGT(FieldFollowerCount, v))
}

// FollowerCountGTE applies the GTE predicate on the "follower_count" field.
func FollowerCountGTE(v int) predicate.UserCount {
	return predicate.UserCount(sql.FieldGTE(FieldFollowerCount, v))
}

// FollowerCountLT applies the LT predicate on the "follower_count" field.
func FollowerCountLT(v int) predicate.UserCount {
	return predicate.UserCount(sql.FieldLT(FieldFollowerCount, v))
}

// FollowerCountLTE applies the LTE predicate on the "follower_count" field.
func FollowerCountLTE(v int) predicate.UserCount {
	return predicate.UserCount(sql.FieldLTE(FieldFollowerCount, v))
}

// FollowingsCountEQ applies the EQ predicate on the "followings_count" field.
func FollowingsCountEQ(v int) predicate.UserCount {
	return predicate.UserCount(sql.FieldEQ(FieldFollowingsCount, v))
}

// FollowingsCountNEQ applies the NEQ predicate on the "followings_count" field.
func FollowingsCountNEQ(v int) predicate.UserCount {
	return predicate.UserCount(sql.FieldNEQ(FieldFollowingsCount, v))
}

// FollowingsCountIn applies the In predicate on the "followings_count" field.
func FollowingsCountIn(vs ...int) predicate.UserCount {
	return predicate.UserCount(sql.FieldIn(FieldFollowingsCount, vs...))
}

// FollowingsCountNotIn applies the NotIn predicate on the "followings_count" field.
func FollowingsCountNotIn(vs ...int) predicate.UserCount {
	return predicate.UserCount(sql.FieldNotIn(FieldFollowingsCount, vs...))
}

// FollowingsCountGT applies the GT predicate on the "followings_count" field.
func FollowingsCountGT(v int) predicate.UserCount {
	return predicate.UserCount(sql.FieldGT(FieldFollowingsCount, v))
}

// FollowingsCountGTE applies the GTE predicate on the "followings_count" field.
func FollowingsCountGTE(v int) predicate.UserCount {
	return predicate.UserCount(sql.FieldGTE(FieldFollowingsCount, v))
}

// FollowingsCountLT applies the LT predicate on the "followings_count" field.
func FollowingsCountLT(v int) predicate.UserCount {
	return predicate.UserCount(sql.FieldLT(FieldFollowingsCount, v))
}

// FollowingsCountLTE applies the LTE predicate on the "followings_count" field.
func FollowingsCountLTE(v int) predicate.UserCount {
	return predicate.UserCount(sql.FieldLTE(FieldFollowingsCount, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.UserCount {
	return predicate.UserCount(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.UserAccount) predicate.UserCount {
	return predicate.UserCount(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserCount) predicate.UserCount {
	return predicate.UserCount(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserCount) predicate.UserCount {
	return predicate.UserCount(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserCount) predicate.UserCount {
	return predicate.UserCount(sql.NotPredicates(p))
}
