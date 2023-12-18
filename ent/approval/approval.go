// Code generated by ent, DO NOT EDIT.

package approval

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the approval type in the database.
	Label = "approval"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPerson holds the string denoting the person field in the database.
	FieldPerson = "person"
	// FieldApprovedTime holds the string denoting the approved_time field in the database.
	FieldApprovedTime = "approved_time"
	// FieldApproved holds the string denoting the approved field in the database.
	FieldApproved = "approved"
	// FieldRevoked holds the string denoting the revoked field in the database.
	FieldRevoked = "revoked"
	// FieldRevokedTime holds the string denoting the revoked_time field in the database.
	FieldRevokedTime = "revoked_time"
	// FieldRequestID holds the string denoting the request_id field in the database.
	FieldRequestID = "request_id"
	// EdgeRequests holds the string denoting the requests edge name in mutations.
	EdgeRequests = "requests"
	// Table holds the table name of the approval in the database.
	Table = "approvals"
	// RequestsTable is the table that holds the requests relation/edge.
	RequestsTable = "approvals"
	// RequestsInverseTable is the table name for the Request entity.
	// It exists in this package in order to avoid circular dependency with the "request" package.
	RequestsInverseTable = "requests"
	// RequestsColumn is the table column denoting the requests relation/edge.
	RequestsColumn = "approval_requests"
)

// Columns holds all SQL columns for approval fields.
var Columns = []string{
	FieldID,
	FieldPerson,
	FieldApprovedTime,
	FieldApproved,
	FieldRevoked,
	FieldRevokedTime,
	FieldRequestID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "approvals"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"approval_requests",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultRevoked holds the default value on creation for the "revoked" field.
	DefaultRevoked bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Approval queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPerson orders the results by the person field.
func ByPerson(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPerson, opts...).ToFunc()
}

// ByApprovedTime orders the results by the approved_time field.
func ByApprovedTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldApprovedTime, opts...).ToFunc()
}

// ByApproved orders the results by the approved field.
func ByApproved(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldApproved, opts...).ToFunc()
}

// ByRevoked orders the results by the revoked field.
func ByRevoked(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRevoked, opts...).ToFunc()
}

// ByRevokedTime orders the results by the revoked_time field.
func ByRevokedTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRevokedTime, opts...).ToFunc()
}

// ByRequestID orders the results by the request_id field.
func ByRequestID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRequestID, opts...).ToFunc()
}

// ByRequestsField orders the results by requests field.
func ByRequestsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRequestsStep(), sql.OrderByField(field, opts...))
	}
}
func newRequestsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RequestsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, RequestsTable, RequestsColumn),
	)
}