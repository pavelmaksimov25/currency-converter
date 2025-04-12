package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pavelmaksimov25/currency-converter/internal/converter"
	"github.com/pavelmaksimov25/currency-converter/internal/exchangerate/api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	exchangeRateApiKey := os.Getenv("EXCHANGE_RATE_API_KEY")
	baseUrl := os.Getenv("EXCHANGE_RATE_API_BASE_URL")

	apiClient := api.NewExchangeRateAPIClient(exchangeRateApiKey, baseUrl)
	converterService := converter.NewConverter(apiClient)

	convertCriteria := converter.ConvertCriteria{
		BaseCurrency:   "USD",
		TargetCurrency: "EUR",
		Amount:         100,
	}

	convertedAmount, err := converterService.Convert(convertCriteria)
	if err != nil {
		log.Fatalf("Error converting currency: %v", err)
	}

	fmt.Printf("Converted amount: %.2f %s\n", convertedAmount, convertCriteria.TargetCurrency)
}
