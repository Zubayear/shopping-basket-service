package mocks

import (
	"ShoppingBasket/ent"
	"context"
	"github.com/google/uuid"
	"gotest.tools/assert"
	"testing"
)

func TestIBasketRepository_AddBasket(t *testing.T) {
	testObj := new(IBasketRepository)
	ctx := context.Background()
	basket := &ent.Basket{
		UserId:     uuid.New(),
		CouponCode: "FREE78",
	}
	expecting := basket
	testObj.On("AddBasket", ctx, basket).Return(expecting, nil)
	addBasket, err := testObj.AddBasket(ctx, basket)
	if err != nil {
		return
	}
	assert.Equal(t, addBasket.CouponCode, basket.CouponCode)
	testObj.AssertExpectations(t)
}

func TestIBasketRepository_GetBasketById(t *testing.T) {
	testObj := new(IBasketRepository)
	ctx := context.Background()
	input := uuid.New()
	expecting := &ent.Basket{
		ID:         uuid.New(),
		UserId:     uuid.New(),
		CouponCode: "FREE90",
	}
	testObj.On("GetBasketById", ctx, input).Return(expecting, nil)
	addBasket, err := testObj.GetBasketById(ctx, input)
	if err != nil {
		return
	}
	assert.Equal(t, addBasket.CouponCode, "FREE90")
	testObj.AssertExpectations(t)
}
