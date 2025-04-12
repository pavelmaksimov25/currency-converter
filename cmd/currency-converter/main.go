package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
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

	rate, err := apiClient.GetExchangeRate("USD", "EUR")
	if err != nil {
		log.Fatalf("Error getting exchange rate: %v", err)
	}

	fmt.Printf("Exchange rate from USD to EUR: %f\n", rate)
}
