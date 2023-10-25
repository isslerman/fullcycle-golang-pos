package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Product 1", 10.0)
	assert.Nil(t, err)  // sem erros
	assert.NotNil(t, p) // product is not null
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "Product 1", p.Name)
	assert.Equal(t, 10.0, p.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 10.0)
	assert.Nil(t, p)                        // our product is nil
	assert.Equal(t, ErrNameIsRequired, err) // we got the right error
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("Product 1", 0)
	assert.Nil(t, p)                         // our product is nil
	assert.Equal(t, ErrPriceIsRequired, err) // we got the right error
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct("Product 1", -10)
	assert.Nil(t, p)                      // our product is nil
	assert.Equal(t, ErrInvalidPrice, err) // we got the right error
}

func TestProductValidate(t *testing.T) {
	p, err := NewProduct("Product 1", 10)
	assert.Nil(t, err)
	assert.NotNil(t, p) // product is not null
	assert.Nil(t, p.Validate())
}
