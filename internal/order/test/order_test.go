package test

import (
	"testing"

	"github.com/josecarlosmoura/gointensivo/internal/order/entity"
	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyID_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := entity.Order{}
	assert.Error(t, order.IsValid(), "invalid id")
}

func TestGivenAnEmptyPrice_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := entity.Order{ID: "123"}
	assert.Error(t, order.IsValid(), "invalid price")
}

func TestGivenAnEmptyTax_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := entity.Order{ID: "123", Price: 10}
	assert.Error(t, order.IsValid(), "invalid tax")
}

func TestGivenAValidParams_WhenICallNewOrder_ThenIShouldReceiveCreateOrderWithAllParams(t *testing.T) {
	order := entity.Order{
		ID:    "123",
		Price: 10.0,
		Tax:   2.0,
	}
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
	assert.Nil(t, order.IsValid())
}

func TestGivenAValidParams_WhenICallNewOrderFunc_ThenIShouldReceiveCreateOrderWithAllParams(t *testing.T) {
	order, err := entity.NewOrder("123", 10.0, 2.0)

	assert.Nil(t, err)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
	assert.Nil(t, order.IsValid())
}

func TestGivenAPriceAndTax_WhenICallCalculateFinalPriceFunc_ThenIShoudSetFinalPrice(t *testing.T) {
	order, err := entity.NewOrder("123", 10.0, 2.0)
	assert.Nil(t, err)
	assert.Nil(t, order.CalculateFinalPrice())
	assert.Equal(t, 12.0, order.FinalPrice)
}
