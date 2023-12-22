// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/orbit-ops/launchpad-core/ent/mission"
	"github.com/orbit-ops/launchpad-core/ent/request"
)

// Request is the model entity for the Request schema.
type Request struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Reason holds the value of the "reason" field.
	Reason string `json:"reason,omitempty"`
	// Requester holds the value of the "requester" field.
	Requester string `json:"requester,omitempty"`
	// RocketConfig holds the value of the "rocket_config" field.
	RocketConfig map[string]string `json:"rocket_config,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RequestQuery when eager-loading is set.
	Edges            RequestEdges `json:"edges"`
	mission_requests *string
	selectValues     sql.SelectValues
}

// RequestEdges holds the relations/edges for other nodes in the graph.
type RequestEdges struct {
	// Mission holds the value of the mission edge.
	Mission *Mission `json:"mission,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// MissionOrErr returns the Mission value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RequestEdges) MissionOrErr() (*Mission, error) {
	if e.loadedTypes[0] {
		if e.Mission == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: mission.Label}
		}
		return e.Mission, nil
	}
	return nil, &NotLoadedError{edge: "mission"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Request) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case request.FieldRocketConfig:
			values[i] = new([]byte)
		case request.FieldReason, request.FieldRequester:
			values[i] = new(sql.NullString)
		case request.FieldID:
			values[i] = new(uuid.UUID)
		case request.ForeignKeys[0]: // mission_requests
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Request fields.
func (r *Request) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case request.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				r.ID = *value
			}
		case request.FieldReason:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reason", values[i])
			} else if value.Valid {
				r.Reason = value.String
			}
		case request.FieldRequester:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field requester", values[i])
			} else if value.Valid {
				r.Requester = value.String
			}
		case request.FieldRocketConfig:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field rocket_config", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &r.RocketConfig); err != nil {
					return fmt.Errorf("unmarshal field rocket_config: %w", err)
				}
			}
		case request.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mission_requests", values[i])
			} else if value.Valid {
				r.mission_requests = new(string)
				*r.mission_requests = value.String
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Request.
// This includes values selected through modifiers, order, etc.
func (r *Request) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryMission queries the "mission" edge of the Request entity.
func (r *Request) QueryMission() *MissionQuery {
	return NewRequestClient(r.config).QueryMission(r)
}

// Update returns a builder for updating this Request.
// Note that you need to call Request.Unwrap() before calling this method if this Request
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Request) Update() *RequestUpdateOne {
	return NewRequestClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Request entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Request) Unwrap() *Request {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Request is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Request) String() string {
	var builder strings.Builder
	builder.WriteString("Request(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("reason=")
	builder.WriteString(r.Reason)
	builder.WriteString(", ")
	builder.WriteString("requester=")
	builder.WriteString(r.Requester)
	builder.WriteString(", ")
	builder.WriteString("rocket_config=")
	builder.WriteString(fmt.Sprintf("%v", r.RocketConfig))
	builder.WriteByte(')')
	return builder.String()
}

// Requests is a parsable slice of Request.
type Requests []*Request
