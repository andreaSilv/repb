package cli

type Option struct {
	Desc    string
	Handler func()
}

var options []*Option
var continueCliLoop bool = true

func CliLoop() {
	NewBlock()
	defer EndBlock()

	for continueCliLoop {
		cli()
	}
}

func RegisterOption(opt *Option) {
	options = append(options, opt)
}

func EndCliLoop() {
	continueCliLoop = false
}

func cli() {
	PrintLn("Choose an option:")

	for i, v := range options {
		Printf("%d - %s\n", i, v.Desc)
	}

	c, err := ReadInt("Choice: ")
	Print("\n")
	if err != nil || c >= len(options) {
		PrintLn("# Invalid value\n")
		return
	}
	Print("\n\n")

	options[c].Handler()
}
