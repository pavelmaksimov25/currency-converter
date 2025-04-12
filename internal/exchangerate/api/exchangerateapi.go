package api

import "github.com/pavelmaksimov25/currency-converter/internal/exchangerate"

type ExchangeRateAPIClient struct {
	apiKey string
}

func NewExchangeRateAPIClient(apiKey string) exchangerate.ExchangeRate {
	return &ExchangeRateAPIClient{
		apiKey: apiKey,
	}
}

func (e * ExchangeRateAPIClient) GetExchangeRate(baseCurrency string, targetCurrency string) (float64, error) {
	// todo :: implement the function to get exchange rate from API
	return 0.0, nil
}
