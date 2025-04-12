package converter

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockExchangeRate is a mock implementation of the ExchangeRate interface
type MockExchangeRate struct {
	mock.Mock
}

func (m *MockExchangeRate) GetExchangeRate(baseCurrency, targetCurrency string) (float64, error) {
	args := m.Called(baseCurrency, targetCurrency)
	return args.Get(0).(float64), args.Error(1)
}

func TestSimpleConverter_Convert_Success(t *testing.T) {
	mockExchangeRate := new(MockExchangeRate)
	mockExchangeRate.On("GetExchangeRate", "USD", "EUR").Return(0.85, nil)

	converter := NewConverter(mockExchangeRate)
	criteria := ConvertCriteria{
		BaseCurrency:   "USD",
		TargetCurrency: "EUR",
		Amount:         100,
	}

	result, err := converter.Convert(criteria)

	assert.NoError(t, err)
	assert.Equal(t, 85.0, result)
	mockExchangeRate.AssertExpectations(t)
}

func TestSimpleConverter_Convert_Error(t *testing.T) {
	mockExchangeRate := new(MockExchangeRate)
	mockExchangeRate.On("GetExchangeRate", "USD", "EUR").Return(0.0, errors.New("exchange rate not found"))

	converter := NewConverter(mockExchangeRate)
	criteria := ConvertCriteria{
		BaseCurrency:   "USD",
		TargetCurrency: "EUR",
		Amount:         100,
	}

	result, err := converter.Convert(criteria)

	assert.Error(t, err)
	assert.Equal(t, 0.0, result)
	mockExchangeRate.AssertExpectations(t)
}
