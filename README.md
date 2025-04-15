# Currency Converter

Real time currency converter

## Core Features

1. Fetching Exchange Rates
    - Use ExchangeRate-API. It has free tier.   
    - Alternative: Open Exchange Rates.

2. Parsing the JSON Response.

3. Performing the Conversion.

## Usage

```bash
go build -o currency-converter ./cmd/currency-converter/main.go
```

```bash
./currency-converter
```

```bash
curl http://localhost:8080/convert?amount=100&base=USD&target=EUR
```

## TODO 
- Introduce mock api client and a flag based on which either real api or mock one will be used.
- Add HTTP interface.
- Use goroutines.
- Cache Exchange Rate API response.
