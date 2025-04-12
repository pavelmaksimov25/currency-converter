package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/pavelmaksimov25/currency-converter/internal/exchangerate"
)

type APIRateResponse struct {
	Success   string             `json:"result"`
	Timestamp int64              `json:"time_last_update_unix"`
	Base      string             `json:"base_code"`
	Date      string             `json:"time_last_update_utc"`
	Rates     map[string]float64 `json:"conversion_rates"`
}

type ExchangeRateAPIClient struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

func NewExchangeRateAPIClient(apiKey string, baseUrl string) exchangerate.ExchangeRate {
	return &ExchangeRateAPIClient{
		apiKey:     apiKey,
		baseURL:    baseUrl,
		httpClient: &http.Client{},
	}
}

func (e *ExchangeRateAPIClient) GetExchangeRate(baseCurrency string, targetCurrency string) (float64, error) {
	// url := e.baseURL + "/" + e.apiKey + "/latest/" + baseCurrency

	// resp, err := e.httpClient.Get(url)
	// if err != nil {
	// 	return 0.0, err
	// }

	// defer resp.Body.Close()

	// if resp.StatusCode != http.StatusOK {
	// 	return 0.0, fmt.Errorf("error: %s", resp.Status)
	// }

	body := e.getMockData() // use mock data to avoid network call during developmment

	var response APIRateResponse
	json.Unmarshal(body, &response)

	rate, ok := response.Rates[targetCurrency]
	if !ok {
		return 0.0, fmt.Errorf("rate for %s not found in response", targetCurrency)
	}

	return rate, nil
}

func (e *ExchangeRateAPIClient) getMockData() []byte {
	data, err := os.ReadFile("./internal/exchangerate/api/mocks/exchange_api_response.json")
	if err != nil {
		panic(err)
	}

	return data
}
