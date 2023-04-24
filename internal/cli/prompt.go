package cli

import (
	"fmt"
	"strconv"
	"strings"
)

var layer int = 0
var layerChar string = "  "

func NewBlock() {
	layer++
}

func EndBlock() {
	layer--
}

func PrintLn(text string) {
	Printf("%s\n", text)
}

func Print(text string) {
	Printf("%s", text)
}

func Printf(format string, vals ...any) {
	formatted := fmt.Sprintf(format, vals...)
	fmt.Printf("%s%s", strings.Repeat(layerChar, layer), formatted)
}

func Read(text string) string {
	Print(text)

	var input string
	fmt.Scanln(&input)
	return input
}

func ReadInt(text string) (int, error) {
	input := Read(text)
	return strconv.Atoi(input)
}
