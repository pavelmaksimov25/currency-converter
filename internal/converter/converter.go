package converter

import (
	"github.com/pavelmaksimov25/currency-converter/internal/exchangerate"
)

type ConvertCriteria struct {
	BaseCurrency   string
	TargetCurrency string
	Amount         float64
}

type Converter interface {
	Convert(convertCriteria *ConvertCriteria) (float64, error)
}

type SimpleConverter struct {
	exchangeRate exchangerate.ExchangeRate
}

func NewConverter(service exchangerate.ExchangeRate) Converter {
	return &SimpleConverter{exchangeRate: service}
}

func (c *SimpleConverter) Convert(convertCriteria *ConvertCriteria) (float64, error) {
	rateData, err := c.exchangeRate.GetExchangeRate(convertCriteria.BaseCurrency, convertCriteria.TargetCurrency)
	if err != nil {
		return 0, err
	}

	return convertCriteria.Amount * rateData, nil
}
