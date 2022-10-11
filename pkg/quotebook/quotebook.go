package quotebook

import (
	"bytes"
	_ "embed"
	"math/rand"
	"time"
)

//go:embed source/*.txt
var source []byte

func init() {
	rand.Seed(time.Now().UnixNano())
}

type QuoteBook [][]byte

func NewQuoteBook() QuoteBook {
	return bytes.Split(source, []byte("\n"))
}

func (QuoteBook QuoteBook) GetRandomQoute() []byte {
	return QuoteBook[rand.Intn(len(QuoteBook))]
}
