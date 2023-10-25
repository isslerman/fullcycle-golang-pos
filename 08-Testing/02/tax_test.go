package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000.0)
	assert.Nil(t, err)
	assert.Equal(t, 10.0, tax)

	tax, err = CalculateTax(0)
	assert.Error(t, err, "amount must be grater than 0")
	assert.Equal(t, 0.0, tax)
	assert.Contains(t, err.Error(), "grater than 0")
}

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	// no caso de passarmos uma tax 10.0, não devemos ter nenhum erro de retorno.
	repository.On("SaveTax", 10.0).Return(nil)
	// repository.On("SaveTax", 10.0).Return(nil).Once() // apenas chamar uma vez
	repository.On("SaveTax", 0.0).Return(errors.New("error saving tax"))
	// repository.On("SaveTax", mock.Anything).Return(errors.New("error saving tax"))

	err := CalculateTaxAndSave(1000.00, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0.0, repository)
	assert.Error(t, err, "error saving tax")

	repository.AssertExpectations(t)
	// repository.AssertCalled(t, "SaveTax", 10.0) // precisa ser chamado 3x
}
