# go-alphavantage

Experimenting with the [Alpha Vantage](https://www.alphavantage.co/documentation/) API by creating a CLI in Go.

## Installation

1. Clone this repository
2. Run `go install` in the clone

## Usage

Set `ALPHAVANTAGE_API_KEY` before you execute the binary or pass the API key as a flag with `-k API_KEY`.

### Main

```
   fx        get exchange rates
   price, p  get share price
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)

```

### FX 

```
USAGE:
   go-alphavantage fx [command options] [FROM SYMBOL] [TO SYMBOL]
```

### Price

```
USAGE:
   go-alphavantage price [command options] [TICKER]
```