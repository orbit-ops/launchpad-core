// Code generated by ent, DO NOT EDIT.

package mission

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the mission type in the database.
	Label = "mission"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldMinApprovers holds the string denoting the min_approvers field in the database.
	FieldMinApprovers = "min_approvers"
	// FieldPossibleApprovers holds the string denoting the possible_approvers field in the database.
	FieldPossibleApprovers = "possible_approvers"
	// EdgeRockets holds the string denoting the rockets edge name in mutations.
	EdgeRockets = "rockets"
	// EdgeRequests holds the string denoting the requests edge name in mutations.
	EdgeRequests = "requests"
	// Table holds the table name of the mission in the database.
	Table = "missions"
	// RocketsTable is the table that holds the rockets relation/edge. The primary key declared below.
	RocketsTable = "rocket_missions"
	// RocketsInverseTable is the table name for the Rocket entity.
	// It exists in this package in order to avoid circular dependency with the "rocket" package.
	RocketsInverseTable = "rockets"
	// RequestsTable is the table that holds the requests relation/edge.
	RequestsTable = "requests"
	// RequestsInverseTable is the table name for the Request entity.
	// It exists in this package in order to avoid circular dependency with the "request" package.
	RequestsInverseTable = "requests"
	// RequestsColumn is the table column denoting the requests relation/edge.
	RequestsColumn = "mission_requests"
)

// Columns holds all SQL columns for mission fields.
var Columns = []string{
	FieldID,
	FieldDescription,
	FieldMinApprovers,
	FieldPossibleApprovers,
}

var (
	// RocketsPrimaryKey and RocketsColumn2 are the table columns denoting the
	// primary key for the rockets relation (M2M).
	RocketsPrimaryKey = []string{"rocket_id", "mission_id"}
)

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
	// MinApproversValidator is a validator for the "min_approvers" field. It is called by the builders before save.
	MinApproversValidator func(int) error
)

// OrderOption defines the ordering options for the Mission queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByMinApprovers orders the results by the min_approvers field.
func ByMinApprovers(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMinApprovers, opts...).ToFunc()
}

// ByRocketsCount orders the results by rockets count.
func ByRocketsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRocketsStep(), opts...)
	}
}

// ByRockets orders the results by rockets terms.
func ByRockets(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRocketsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByRequestsCount orders the results by requests count.
func ByRequestsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRequestsStep(), opts...)
	}
}

// ByRequests orders the results by requests terms.
func ByRequests(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRequestsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newRocketsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RocketsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, RocketsTable, RocketsPrimaryKey...),
	)
}
func newRequestsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RequestsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RequestsTable, RequestsColumn),
	)
}
