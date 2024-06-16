package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wendellnd/graduate-go-expert-classes/api/pkg/entity"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Product 1", 100)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, 100, product.Price)
	assert.NotEmpty(t, product.CreatedAt)
}

func TestProductTestProductWhenNameIsRequired(t *testing.T) {
	product, err := NewProduct("", 100)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	product, err := NewProduct("Product 1", 0)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	product, err := NewProduct("Product 1", -100)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestProductValidate(t *testing.T) {
	product := Product{
		ID:    entity.NewID(),
		Name:  "Product 1",
		Price: 100,
	}
	err := product.Validate()
	assert.Nil(t, err)
}
