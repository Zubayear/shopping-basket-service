// Code generated by entc, DO NOT EDIT.

package ent

import (
	"ShoppingBasket/ent/basket"
	"ShoppingBasket/ent/basketline"
	"ShoppingBasket/ent/event"
	"ShoppingBasket/ent/schema"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	basketFields := schema.Basket{}.Fields()
	_ = basketFields
	// basketDescCouponCode is the schema descriptor for CouponCode field.
	basketDescCouponCode := basketFields[2].Descriptor()
	// basket.CouponCodeValidator is a validator for the "CouponCode" field. It is called by the builders before save.
	basket.CouponCodeValidator = basketDescCouponCode.Validators[0].(func(string) error)
	// basketDescID is the schema descriptor for id field.
	basketDescID := basketFields[0].Descriptor()
	// basket.DefaultID holds the default value on creation for the id field.
	basket.DefaultID = basketDescID.Default.(func() uuid.UUID)
	basketlineFields := schema.BasketLine{}.Fields()
	_ = basketlineFields
	// basketlineDescID is the schema descriptor for id field.
	basketlineDescID := basketlineFields[0].Descriptor()
	// basketline.DefaultID holds the default value on creation for the id field.
	basketline.DefaultID = basketlineDescID.Default.(func() uuid.UUID)
	eventFields := schema.Event{}.Fields()
	_ = eventFields
	// eventDescDate is the schema descriptor for Date field.
	eventDescDate := eventFields[2].Descriptor()
	// event.DefaultDate holds the default value on creation for the Date field.
	event.DefaultDate = eventDescDate.Default.(func() time.Time)
	// event.UpdateDefaultDate holds the default value on update for the Date field.
	event.UpdateDefaultDate = eventDescDate.UpdateDefault.(func() time.Time)
	// eventDescID is the schema descriptor for id field.
	eventDescID := eventFields[0].Descriptor()
	// event.DefaultID holds the default value on creation for the id field.
	event.DefaultID = eventDescID.Default.(func() uuid.UUID)
}
