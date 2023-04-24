package main

import (
	"github.com/andreaSilv/repb/internal/cli"
	_ "github.com/andreaSilv/repb/internal/option"
)

func main() {
	cli.PrintLn("## Welcome to repb ##")
	cli.CliLoop()
}
