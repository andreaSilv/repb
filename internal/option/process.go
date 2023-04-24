package option

import (
	"github.com/andreaSilv/repb/internal/cli"
	"github.com/andreaSilv/repb/internal/proc"
)

func processHandler() {
	var transactions []proc.Transaction

	for _, i := range proc.StoredInputs {
		transactions = append(transactions, applyAllParsers(i)...)
	}

	proc.StoredTransactions = transactions
}

func applyAllParsers(inString string) []proc.Transaction {
	parsers := [...]proc.TransactionParser{
		proc.NewAmexTransactionParser(&inString),
		proc.NewFinecoTransactionParser(&inString),
	}

	var transactions []proc.Transaction
	for _, i := range parsers {
		transactions = append(transactions, i.Parse()...)
	}

	return transactions
}

func init() {
	cli.RegisterOption(&cli.Option{
		Desc:    "Process",
		Handler: processHandler,
	})
}
