package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/pavelmaksimov25/currency-converter/internal/converter"
	"github.com/pavelmaksimov25/currency-converter/internal/exchangerate/api"
	internalHttp "github.com/pavelmaksimov25/currency-converter/internal/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	exchangeRateApiKey := os.Getenv("EXCHANGE_RATE_API_KEY")
	baseUrl := os.Getenv("EXCHANGE_RATE_API_BASE_URL")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	apiClient := api.NewExchangeRateAPIClient(exchangeRateApiKey, baseUrl)
	converterService := converter.NewConverter(apiClient)

	httpHandler := internalHttp.NewHandler(converterService)
	http.Handle("/convert", httpHandler)

	fmt.Printf("Server listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
