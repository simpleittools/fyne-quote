package assets

import (
	_ "embed"
)

//go:embed quotesData.json
var QuotesData []byte

func GetQuotesData() []byte {
	return QuotesData
}
