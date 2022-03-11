package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Basket holds the schema definition for the Basket entity.
type Basket struct {
	ent.Schema
}

// Fields of the Basket.
func (Basket) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("UserId", uuid.UUID{}),
		field.Text("CouponCode").MaxLen(6).Unique(),
	}
}

// Edges of the Basket.
func (Basket) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("BasketLine", BasketLine.Type),
	}
}
