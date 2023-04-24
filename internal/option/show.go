package option

import (
	"fmt"

	"github.com/andreaSilv/repb/internal/cli"
	"github.com/andreaSilv/repb/internal/proc"
	"github.com/atotto/clipboard"
)

func showRaw() {
	cli.PrintLn("Stored:")

	for i, v := range proc.StoredInputs {
		cli.Printf("#%d - %s\n", i, v)
	}
}

func showParsed() {
	cli.PrintLn("Stored:")

	for _, v := range proc.StoredTransactions {
		cli.Printf("%d\t%s\t%s\t%.2f\n", v.DayOfTheMonth, v.Desc, v.Currency, v.Amount)
	}
}

func showParsedInClipboard() {
	var outString string

	for _, v := range proc.StoredTransactions {
		outString += fmt.Sprintf("%d\t%s\t%s\t%.2f\n", v.DayOfTheMonth, v.Desc, v.Currency, v.Amount)
	}

	clipboard.WriteAll(outString)

	cli.PrintLn("Put in your clipboard")
}

func init() {
	cli.RegisterOption(&cli.Option{Desc: "Show raw", Handler: showRaw})
	cli.RegisterOption(&cli.Option{Desc: "Show parsed", Handler: showParsed})
	cli.RegisterOption(&cli.Option{Desc: "Put parsed in clipboard", Handler: showParsedInClipboard})
}
