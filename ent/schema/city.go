package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// City holds the schema definition for the City entity.
type City struct {
	ent.Schema
}

// Edges of the City.
func (City) Edges() []ent.Edge {
	return nil
}

// Fields of the City.
func (City) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").MaxLen(35).Optional(),
		field.String("CountryCode").MaxLen(3).Optional(),
		field.String("District").MaxLen(20).Optional(),
		field.Int("Population").Optional(),
	}
}
