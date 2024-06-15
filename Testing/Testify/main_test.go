package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	amount := 500.0

	result := CalculateTax(amount)
	assert.Equal(t, 5.0, result)
}

func TestCalculateTestAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil)
	//repository.On("SaveTax", 10.0).Return(nil).Once()
	//repository.On("SaveTax", 10.0).Return(nil).Twice()
	repository.On("SaveTax", 0.0).Return(errors.New("error saving tax"))
	//repository.On("SaveTax", mock.Anything).Return(errors.New("error saving tax"))

	err := CalculateTaxAndSave(1000.00, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0.0, repository)
	assert.Error(t, err, "error saving tax")

	repository.AssertExpectations(t)
}
