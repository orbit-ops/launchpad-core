// Code generated by ent, DO NOT EDIT.

package apikey

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the apikey type in the database.
	Label = "api_key"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// Table holds the table name of the apikey in the database.
	Table = "api_keys"
)

// Columns holds all SQL columns for apikey fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldKey,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// OrderOption defines the ordering options for the ApiKey queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByKey orders the results by the key field.
func ByKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKey, opts...).ToFunc()
}
