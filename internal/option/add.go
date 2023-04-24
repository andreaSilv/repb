package option

import (
	"github.com/andreaSilv/repb/internal/cli"
	"github.com/andreaSilv/repb/internal/proc"
	"github.com/atotto/clipboard"

)

func addString() {
	val := cli.Read("Paste here your table: ")
	proc.StoredInputs = append(proc.StoredInputs, val)
}

func addClipboard() {
	val, _ := clipboard.ReadAll()

	proc.StoredInputs = append(proc.StoredInputs, val)
	cli.PrintLn("Added!")
}

func init() {
	cli.RegisterOption(&cli.Option{Desc: "Add string", Handler: addString})
	cli.RegisterOption(&cli.Option{Desc: "Add clipboard", Handler: addClipboard})
}
