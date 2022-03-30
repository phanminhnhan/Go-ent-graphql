package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
	"time"
	"github.com/google/uuid"
)

// Subject holds the schema definition for the Subject entity.
type Subject struct {
	ent.Schema
}

// Fields of the Subject.
func (Subject) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("subject_id", uuid.UUID{}).
            Default(uuid.New),
		field.String("name").
            Default("unknown"),
		field.Time("created_at").
            Default(time.Now),	
		field.Time("updated_at").
			Optional().
            Nillable(),		
	}
}

// Edges of the Subject.
func (Subject) Edges() []ent.Edge {
	return []ent.Edge{
        edge.From("user", User.Type).
            Ref("subjects").
            Unique(),
    }
}
