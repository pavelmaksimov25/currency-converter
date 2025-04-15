package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pavelmaksimov25/currency-converter/internal/converter"
)

type Handler struct {
	converter converter.Converter
}

func NewHandler(converter converter.Converter) *Handler {
	return &Handler{
		converter: converter,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	amountStr := r.URL.Query().Get("amount")
	baseCurrency := r.URL.Query().Get("base")
	targetCurrency := r.URL.Query().Get("target")

	if amountStr == "" || baseCurrency == "" || targetCurrency == "" {
		http.Error(w, "Missing required parameters: amount, base, or target", http.StatusBadRequest)
		return
	}

	// todo :: valdidate better

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	criteria := &converter.ConvertCriteria{
		BaseCurrency:   baseCurrency,
		TargetCurrency: targetCurrency,
		Amount:         amount,
	}

	result, err := h.converter.Convert(criteria)
	if err != nil {
		http.Error(w, "Failed to perform conversion: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"amount":          amount,
		"base_currency":   baseCurrency,
		"target_currency": targetCurrency,
		"result":          result,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
