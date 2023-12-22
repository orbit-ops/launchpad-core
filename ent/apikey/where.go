// Code generated by ent, DO NOT EDIT.

package apikey

import (
	"entgo.io/ent/dialect/sql"
	"github.com/orbit-ops/launchpad-core/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldEQ(FieldName, v))
}

// Key applies equality check predicate on the "key" field. It's identical to KeyEQ.
func Key(v string) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldEQ(FieldKey, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldContainsFold(FieldName, v))
}

// KeyEQ applies the EQ predicate on the "key" field.
func KeyEQ(v string) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldEQ(FieldKey, v))
}

// KeyNEQ applies the NEQ predicate on the "key" field.
func KeyNEQ(v string) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldNEQ(FieldKey, v))
}

// KeyIn applies the In predicate on the "key" field.
func KeyIn(vs ...string) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldIn(FieldKey, vs...))
}

// KeyNotIn applies the NotIn predicate on the "key" field.
func KeyNotIn(vs ...string) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldNotIn(FieldKey, vs...))
}

// KeyGT applies the GT predicate on the "key" field.
func KeyGT(v string) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldGT(FieldKey, v))
}

// KeyGTE applies the GTE predicate on the "key" field.
func KeyGTE(v string) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldGTE(FieldKey, v))
}

// KeyLT applies the LT predicate on the "key" field.
func KeyLT(v string) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldLT(FieldKey, v))
}

// KeyLTE applies the LTE predicate on the "key" field.
func KeyLTE(v string) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldLTE(FieldKey, v))
}

// KeyContains applies the Contains predicate on the "key" field.
func KeyContains(v string) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldContains(FieldKey, v))
}

// KeyHasPrefix applies the HasPrefix predicate on the "key" field.
func KeyHasPrefix(v string) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldHasPrefix(FieldKey, v))
}

// KeyHasSuffix applies the HasSuffix predicate on the "key" field.
func KeyHasSuffix(v string) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldHasSuffix(FieldKey, v))
}

// KeyEqualFold applies the EqualFold predicate on the "key" field.
func KeyEqualFold(v string) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldEqualFold(FieldKey, v))
}

// KeyContainsFold applies the ContainsFold predicate on the "key" field.
func KeyContainsFold(v string) predicate.ApiKey {
	return predicate.ApiKey(sql.FieldContainsFold(FieldKey, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ApiKey) predicate.ApiKey {
	return predicate.ApiKey(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ApiKey) predicate.ApiKey {
	return predicate.ApiKey(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ApiKey) predicate.ApiKey {
	return predicate.ApiKey(sql.NotPredicates(p))
}
