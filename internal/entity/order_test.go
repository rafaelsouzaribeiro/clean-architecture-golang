package entity

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyId(t *testing.T) {
	order := Order{}
	assert.Equal(t, order.isValid(), errors.New("invalid id"))
}

func TestGivenAnEmptyPrice(t *testing.T) {
	order := Order{ID: "aaa"}
	assert.Equal(t, order.isValid(), errors.New("invalid price"))
}

func TestGivenAnEmptyTax(t *testing.T) {
	order := Order{ID: "aaa", Price: 10}
	assert.Equal(t, order.isValid(), errors.New("invalid tax"))
}

func TestGivenAValidParams(t *testing.T) {
	order := Order{
		ID:    "aaa",
		Price: 10.0,
		Tax:   5.0}
	assert.Equal(t, "aaa", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 5.0, order.Tax)
	assert.Nil(t, order.isValid())
}

func TestGivenParamsCallNewOrder(t *testing.T) {
	order, err := NewOrder("aaa", 10.0, 5.0)
	assert.Equal(t, "aaa", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 5.0, order.Tax)
	assert.Nil(t, err)
}

func TestGivenParamsCallNewOrderAndCalculate(t *testing.T) {
	order, err := NewOrder("aaa", 10.0, 5.0)
	assert.Nil(t, err)
	assert.Nil(t, order.CalculateFinalPrice())
	assert.Equal(t, order.FinalPrice, 10.5)

}
