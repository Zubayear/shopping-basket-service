package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BasketLine holds the schema definition for the BasketLine entity.
type BasketLine struct {
	ent.Schema
}

// Fields of the BasketLine.
func (BasketLine) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Uint("TicketAmount"),
		field.Float32("Price").
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(6,2)",
			}),
	}
}

// Edges of the BasketLine.
func (BasketLine) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Event", Event.Type).Unique(),
		edge.From("Basket", Basket.Type).Ref("BasketLine").Unique(),
	}
}
