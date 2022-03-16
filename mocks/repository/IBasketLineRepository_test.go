package mocks

import (
	"ShoppingBasket/ent"
	"context"
	"github.com/google/uuid"
	"gotest.tools/assert"
	"testing"
)

func TestIBasketLineRepository_AddBasketLine(t *testing.T) {
	testObj := new(IBasketLineRepository)
	ctx := context.Background()
	basketId := uuid.New()
	basketLine := &ent.BasketLine{
		ID:           uuid.New(),
		TicketAmount: 7,
		Price:        890.90,
	}
	expecting := basketLine
	testObj.On("AddBasketLine", ctx, basketId, basketLine).Return(expecting, nil)
	testObj.AddBasketLine(ctx, basketId, basketLine)
	testObj.AssertExpectations(t)
}

func TestIBasketLineRepository_DeleteBasketLine(t *testing.T) {
	testObj := new(IBasketLineRepository)
	ctx := context.Background()
	basketLineId := uuid.New()
	expecting := "Basket Line Deleted"
	testObj.On("DeleteBasketLine", ctx, basketLineId).Return(expecting, nil)
	msg, err := testObj.DeleteBasketLine(ctx, basketLineId)
	if err != nil {
		return
	}
	assert.Equal(t, expecting, msg)
	testObj.AssertExpectations(t)
}

func TestIBasketLineRepository_GetBasketLineById(t *testing.T) {
	testObj := new(IBasketLineRepository)
	ctx := context.Background()
	basketLineId := uuid.New()
	expecting := &ent.BasketLine{
		ID:           uuid.New(),
		TicketAmount: 7,
		Price:        890.90,
	}
	testObj.On("GetBasketLineById", ctx, basketLineId).Return(expecting, nil)
	testObj.GetBasketLineById(ctx, basketLineId)
	testObj.AssertExpectations(t)
}

func TestIBasketLineRepository_GetBasketLinesByBasketId(t *testing.T) {
	testObj := new(IBasketLineRepository)
	ctx := context.Background()
	basketId := uuid.New()
	basketLine := &ent.BasketLine{
		ID:           uuid.New(),
		TicketAmount: 7,
		Price:        890.90,
	}
	var expecting []*ent.BasketLine
	expecting = append(expecting, basketLine)
	testObj.On("GetBasketLinesByBasketId", ctx, basketId).Return(expecting, nil)
	basketLines, err := testObj.GetBasketLinesByBasketId(ctx, basketId)
	if err != nil {
		return
	}
	assert.Equal(t, basketLine.Price, basketLines[0].Price)
	testObj.AssertExpectations(t)
}

func TestIBasketLineRepository_UpdateBasketLine(t *testing.T) {
	testObj := new(IBasketLineRepository)
	ctx := context.Background()
	basketLine := &ent.BasketLine{
		ID:           uuid.New(),
		TicketAmount: 7,
		Price:        890.90,
	}
	expecting := basketLine
	testObj.On("UpdateBasketLine", ctx, basketLine).Return(expecting, nil)
	testObj.UpdateBasketLine(ctx, basketLine)
	testObj.AssertExpectations(t)
}
