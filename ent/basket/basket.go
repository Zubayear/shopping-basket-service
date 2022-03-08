// Code generated by entc, DO NOT EDIT.

package basket

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the basket type in the database.
	Label = "basket"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserId holds the string denoting the userid field in the database.
	FieldUserId = "user_id"
	// EdgeBasketLine holds the string denoting the basketline edge name in mutations.
	EdgeBasketLine = "BasketLine"
	// Table holds the table name of the basket in the database.
	Table = "baskets"
	// BasketLineTable is the table that holds the BasketLine relation/edge.
	BasketLineTable = "basket_lines"
	// BasketLineInverseTable is the table name for the BasketLine entity.
	// It exists in this package in order to avoid circular dependency with the "basketline" package.
	BasketLineInverseTable = "basket_lines"
	// BasketLineColumn is the table column denoting the BasketLine relation/edge.
	BasketLineColumn = "basket_basket_line"
)

// Columns holds all SQL columns for basket fields.
var Columns = []string{
	FieldID,
	FieldUserId,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
