package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
	"time"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
            Default(uuid.New),
		field.String("name").
            Default("unknown"),
		field.String("email").
			Unique(),
		field.String("password").
        	Default("unknown"),		
        field.Int("age").
            Positive(),
		field.String("avatar").
			Optional().
			Nillable(),
		field.Time("created_at").
            Default(time.Now),	
		field.Time("updated_at").
			Optional().
            Nillable(),					
    }
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
