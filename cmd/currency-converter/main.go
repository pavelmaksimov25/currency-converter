package main

import (
	"flag"
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

	amountPtr := flag.Float64("amount", 1.0, "The amount to convert")
	basePtr := flag.String("base", "USD", "The base currency code")
	targetPtr := flag.String("target", "EUR", "The target currency code")
	flag.Parse()

	convertCriteria := converter.ConvertCriteria{
		BaseCurrency:   *basePtr,
		TargetCurrency: *targetPtr,
		Amount:         *amountPtr,
	}

	fmt.Printf("Converting %.2f %s to %s...\n", *amountPtr, *basePtr, *targetPtr)

	convertedAmount, err := converterService.Convert(convertCriteria)
	if err != nil {
		log.Fatalf("Error converting currency: %v", err)
	}

	fmt.Printf("Converted amount: %.2f %s\n", convertedAmount, convertCriteria.TargetCurrency)
}
