// Code generated by entc, DO NOT EDIT.

package basketline

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the basketline type in the database.
	Label = "basket_line"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTicketAmount holds the string denoting the ticketamount field in the database.
	FieldTicketAmount = "ticket_amount"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// EdgeEvent holds the string denoting the event edge name in mutations.
	EdgeEvent = "Event"
	// EdgeBasket holds the string denoting the basket edge name in mutations.
	EdgeBasket = "Basket"
	// Table holds the table name of the basketline in the database.
	Table = "basket_lines"
	// EventTable is the table that holds the Event relation/edge.
	EventTable = "events"
	// EventInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	EventInverseTable = "events"
	// EventColumn is the table column denoting the Event relation/edge.
	EventColumn = "basket_line_event"
	// BasketTable is the table that holds the Basket relation/edge.
	BasketTable = "basket_lines"
	// BasketInverseTable is the table name for the Basket entity.
	// It exists in this package in order to avoid circular dependency with the "basket" package.
	BasketInverseTable = "baskets"
	// BasketColumn is the table column denoting the Basket relation/edge.
	BasketColumn = "basket_basket_line"
)

// Columns holds all SQL columns for basketline fields.
var Columns = []string{
	FieldID,
	FieldTicketAmount,
	FieldPrice,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "basket_lines"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"basket_basket_line",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
