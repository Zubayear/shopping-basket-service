package mocks

import (
	"ShoppingBasket/ent"
	"context"
	"github.com/google/uuid"
	"gotest.tools/assert"
	"testing"
	"time"
)

func TestIEventRepository_AddEvent(t *testing.T) {
	testObj := new(IEventRepository)
	ctx := context.Background()
	event := &ent.Event{
		ID:   uuid.New(),
		Name: "Interstellar",
		Date: time.Now(),
	}
	expecting := event
	testObj.On("AddEvent", ctx, event).Return(expecting, nil)
	addedEvent, err := testObj.AddEvent(ctx, event)
	if err != nil {
		return
	}
	assert.Equal(t, addedEvent.Date, event.Date)
	testObj.AssertExpectations(t)
}

func TestIEventRepository_GetEventById(t *testing.T) {
	testObj := new(IEventRepository)
	ctx := context.Background()
	input := uuid.New()
	event := &ent.Event{
		ID:   uuid.New(),
		Name: "Interstellar",
		Date: time.Now(),
	}
	expecting := event
	testObj.On("GetEventById", ctx, input).Return(expecting, nil)
	testObj.GetEventById(ctx, input)
	testObj.AssertExpectations(t)
}
