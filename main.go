package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	app := &cli.App{
		Name:                 "go-alphavantage",
		Description:          "playground for AlphaVantage API",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:      "fx",
				Usage:     "get exchange rates",
				ArgsUsage: "[FROM SYMBOL] [TO SYMBOL]",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "apiKey",
						Aliases: []string{"k"},
						Usage:   "API key for AlphaVantage",
						EnvVars: []string{"ALPHAVANTAGE_API_KEY", "API_KEY"},
					},
				},
				Action: fx,
			},
			{
				Name:      "price",
				Aliases:   []string{"p"},
				Usage:     "get share price",
				ArgsUsage: "[TICKER]",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "apiKey",
						Aliases: []string{"k"},
						Usage:   "API key for AlphaVantage",
						EnvVars: []string{"ALPHAVANTAGE_API_KEY", "API_KEY"},
					},
				},
				Action: price,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

type currencyExchangeResponse struct {
	RealtimeExchangeRate struct {
		ExchangeRate string `json:"5. Exchange Rate"`
	} `json:"Realtime Currency Exchange Rate"`
}
type globalQuoteResponse struct {
	GlobalQuote struct {
		Price string `json:"05. price"`
	} `json:"Global Quote"`
}

func alphaVantageQuery(function string, apiKey string, params url.Values) (*http.Response, error) {
	params.Set("function", function)
	params.Set("apikey", apiKey)

	resp, err := http.Get(fmt.Sprintf("https://alphavantage.co/query?%s", params.Encode()))
	if err != nil {
		return nil, errors.Wrap(err, "could not query AlphaVantage")
	}
	return resp, nil
}

func fx(cCtx *cli.Context) error {
	fromCurrency := cCtx.Args().Get(0)
	toCurrency := cCtx.Args().Get(1)

	params := url.Values{}
	params.Set("from_currency", fromCurrency)
	params.Set("to_currency", toCurrency)

	resp, err := alphaVantageQuery("CURRENCY_EXCHANGE_RATE", cCtx.String("apiKey"), params)
	if err != nil {
		return errors.Wrap(err, "could not execute CURRENCY_EXCHANGE_RATE AlphaVantage query")
	}

	var respData currencyExchangeResponse
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return errors.Wrap(err, "could not decode json")
	}

	fmt.Printf("FX %v/%v: %v\n", fromCurrency, toCurrency, respData.RealtimeExchangeRate.ExchangeRate)

	return nil
}

func price(cCtx *cli.Context) error {
	symbol := cCtx.Args().Get(0)

	params := url.Values{}
	params.Set("symbol", symbol)

	resp, err := alphaVantageQuery("GLOBAL_QUOTE", cCtx.String("apiKey"), params)
	if err != nil {
		return errors.Wrap(err, "could not execute GLOBAL_QUOTE AlphaVantage query")
	}

	var respData globalQuoteResponse
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return errors.Wrap(err, "could not decode json")
	}
	fmt.Printf("Price %v: %v\n", symbol, respData.GlobalQuote.Price)

	return nil
}
