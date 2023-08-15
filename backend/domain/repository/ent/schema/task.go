package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("name"),
		field.String("description").Optional().Nillable(),
		field.Time("deadline").Optional().Nillable(),
		field.Time("completd_at").Optional().Nillable(),
		field.Time("created_at"),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tags", Tag.Type),
	}
}
