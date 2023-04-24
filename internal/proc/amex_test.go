package proc

import (
	"fmt"
	"testing"
)

func TestAmexTransactionParser(t *testing.T) {
	ex1 := `
24 apr
In sospeso
AMAZON IT MARKETPLACE APE

34,00 €

23 apr
In sospeso
ESSO

103,00 €

23 apr
In sospeso
TRATTORIA PANE E OLIO

62,00 €

22 apr
In sospeso
MEDICAL ORTOPEDIA

25,00 €

19 apr
Credito

ADDEBITO IN C/C SALVO BUON FINE

-516,59 €

17 apr
RISTORANTE A FIGLIA DO NAPOLI

87,00 €

17 apr
UNICOCAMPANIA NAPOLI

3,60 €

16 apr
MUSA SRL BACOLI

20,00 €`

	tp := NewAmexTransactionParser(&ex1)
	transactions := tp.Parse()

	sampleOut := ""
	for _, i := range transactions {
		sampleOut += fmt.Sprintf("%d;%s;%s;%.2f\n", i.DayOfTheMonth, i.Desc, i.Currency, i.Amount)
	}

	expected := `24;AMAZON IT MARKETPLACE APE;E;34.00
23;ESSO;E;103.00
23;TRATTORIA PANE E OLIO;E;62.00
22;MEDICAL ORTOPEDIA;E;25.00
17;RISTORANTE A FIGLIA DO NAPOLI;E;87.00
17;UNICOCAMPANIA NAPOLI;E;3.60
16;MUSA SRL BACOLI;E;20.00
`

	if expected != sampleOut {
		t.Error("Unexpected result")
	}
}
