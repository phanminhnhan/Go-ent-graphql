// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"go-ent-graphql/ent/subject"
	"go-ent-graphql/ent/user"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Subject is the model entity for the Subject schema.
type Subject struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// SubjectID holds the value of the "subject_id" field.
	SubjectID uuid.UUID `json:"subject_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SubjectQuery when eager-loading is set.
	Edges         SubjectEdges `json:"edges"`
	user_subjects *int
}

// SubjectEdges holds the relations/edges for other nodes in the graph.
type SubjectEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SubjectEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Subject) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case subject.FieldID:
			values[i] = new(sql.NullInt64)
		case subject.FieldName:
			values[i] = new(sql.NullString)
		case subject.FieldCreatedAt, subject.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case subject.FieldSubjectID:
			values[i] = new(uuid.UUID)
		case subject.ForeignKeys[0]: // user_subjects
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Subject", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Subject fields.
func (s *Subject) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case subject.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case subject.FieldSubjectID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field subject_id", values[i])
			} else if value != nil {
				s.SubjectID = *value
			}
		case subject.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case subject.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case subject.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = new(time.Time)
				*s.UpdatedAt = value.Time
			}
		case subject.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_subjects", value)
			} else if value.Valid {
				s.user_subjects = new(int)
				*s.user_subjects = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Subject entity.
func (s *Subject) QueryUser() *UserQuery {
	return (&SubjectClient{config: s.config}).QueryUser(s)
}

// Update returns a builder for updating this Subject.
// Note that you need to call Subject.Unwrap() before calling this method if this Subject
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Subject) Update() *SubjectUpdateOne {
	return (&SubjectClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Subject entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Subject) Unwrap() *Subject {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Subject is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Subject) String() string {
	var builder strings.Builder
	builder.WriteString("Subject(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", subject_id=")
	builder.WriteString(fmt.Sprintf("%v", s.SubjectID))
	builder.WriteString(", name=")
	builder.WriteString(s.Name)
	builder.WriteString(", created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	if v := s.UpdatedAt; v != nil {
		builder.WriteString(", updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Subjects is a parsable slice of Subject.
type Subjects []*Subject

func (s Subjects) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
