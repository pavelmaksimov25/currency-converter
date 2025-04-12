package exchangerate

type ExchangeRate interface {
	GetExchangeRate(baseCurrency string, targetCurrency string) (float64, error)
}
