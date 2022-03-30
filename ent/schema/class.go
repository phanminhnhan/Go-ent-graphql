package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
	"time"
	"github.com/google/uuid"
)
// Class holds the schema definition for the Class entity.
type Class struct {
	ent.Schema
}

// Fields of the Class.
func (Class) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("class_id", uuid.UUID{}).
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

// Edges of the Class.
func (Class) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
	}
}
