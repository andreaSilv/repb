package option

import "github.com/andreaSilv/repb/internal/cli"

func exitHandler() {
	cli.EndCliLoop()
}

func init() {
	cli.RegisterOption(&cli.Option{
		Desc:    "Exit",
		Handler: exitHandler,
	})
}
