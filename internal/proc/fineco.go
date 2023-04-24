package proc

import (
	"regexp"
	"strconv"
	"strings"
)

type finecoTransactionParser struct {
	raw *string
}

func (c finecoTransactionParser) Parse() []Transaction {
	var transactions []Transaction

	currencyMapping := map[string]string{
		"EUR": EUR,
		"USD": USD,
	}

	pat := regexp.MustCompile(`(\d{1,2})\n([a-z]{1,})\n(.+)\n.+\n-?(\d+,\d+) (...)`)
	matches := pat.FindAllStringSubmatch(*c.raw, -1)

	for _, match := range matches {
		amount, _ := strconv.ParseFloat(strings.Replace(match[4], ",", ".", 1), 32)
		dayOfTheMonth, _ := strconv.ParseInt(match[1], 10, 32)

		transactions = append(transactions, Transaction{
			DayOfTheMonth: int32(dayOfTheMonth),
			Desc:          match[3],
			Amount:        float32(amount),
			Currency:      currencyMapping[match[5]],
		})
	}

	return transactions
}

func NewFinecoTransactionParser(transactionsRaw *string) finecoTransactionParser {
	return finecoTransactionParser{raw: transactionsRaw}
}
