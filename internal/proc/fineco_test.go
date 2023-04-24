package proc

import (
	"fmt"
	"testing"
)

func TestFinecoTransactionParser(t *testing.T) {
	ex1 := `Aprile 2023
23
aprile
Finecopay Da Raffaella Bianconi - Fineco Pay
Finecopay
25,00 EUR
21
aprile
Prelevamento Carta N° ***** 341 Data Operazione 20/4/2023 Ora 19:37 3069 Intesa Sanpaolo Spa N° Atm 1984 Albano Laziale (Rm)
Prelievo Bancomat
-250,00 EUR
19
aprile
American Express Italia Srl Addebito Sdd Fattura A Vs Carico Da It07aex0000014445281000 Mand E2022101065519445 Per 71003
Sepa Direct Debit
-516,59 EUR
18
aprile
La Paranza Societa Coo Napoli It Carta N. ***** 412 Data Operazione 16/04/23
Pagamento Visa Debit
-3,20 EUR
18
aprile
Terrazza Nazionale Napoli It Carta N. ***** 412 Data Operazione 15/04/23
Pagamento Visa Debit
-10,00 EUR
16
aprile
Pag. Del 15/04/23 Ora 18:28 Presso: Poppella Via Santa Brigida 69/70 Napoli 80132 Ita Carta N° *****341 Nessuna Commissione
Pagobancomat Pos
-7,50 EUR
13
aprile
Pag. Del 12/04/23 Ora 09:34 Presso: Pv5649 Via Ardeatina 972 Roma 00134 Ita Carta N° *****341 Nessuna Commissione
Pagobancomat Pos
-30,00 EUR
13
aprile
Punto Verde S.R.L. Di. Castel Gandol It Carta N. ***** 412 Data Operazione 11/04/23
Pagamento Visa Debit
-69,00 EUR
12
aprile
Pag. Del 11/04/23 Ora 19:07 Presso: Bricostore Via Nettunanse Ariccia 00072 Ita Carta N° *****341 Nessuna Commissione
Pagobancomat Pos
-41,58 EUR
09
aprile
Pag. Del 08/04/23 Ora 23:11 Presso: Sam Beer Viale Dell'aeronautica, Roma 00144 Ita Carta N° *****341 Nessuna Commissione
Pagobancomat Pos
-5,00 EUR`

	tp := NewFinecoTransactionParser(&ex1)
	transactions := tp.Parse()

	sampleOut := ""
	for _, i := range transactions {
		sampleOut += fmt.Sprintf("%d;%s;%s;%.2f\n", i.DayOfTheMonth, i.Desc, i.Currency, i.Amount)
	}

	expected := `23;Finecopay Da Raffaella Bianconi - Fineco Pay;E;25.00
21;Prelevamento Carta N° ***** 341 Data Operazione 20/4/2023 Ora 19:37 3069 Intesa Sanpaolo Spa N° Atm 1984 Albano Laziale (Rm);E;250.00
19;American Express Italia Srl Addebito Sdd Fattura A Vs Carico Da It07aex0000014445281000 Mand E2022101065519445 Per 71003;E;516.59
18;La Paranza Societa Coo Napoli It Carta N. ***** 412 Data Operazione 16/04/23;E;3.20
18;Terrazza Nazionale Napoli It Carta N. ***** 412 Data Operazione 15/04/23;E;10.00
16;Pag. Del 15/04/23 Ora 18:28 Presso: Poppella Via Santa Brigida 69/70 Napoli 80132 Ita Carta N° *****341 Nessuna Commissione;E;7.50
13;Pag. Del 12/04/23 Ora 09:34 Presso: Pv5649 Via Ardeatina 972 Roma 00134 Ita Carta N° *****341 Nessuna Commissione;E;30.00
13;Punto Verde S.R.L. Di. Castel Gandol It Carta N. ***** 412 Data Operazione 11/04/23;E;69.00
12;Pag. Del 11/04/23 Ora 19:07 Presso: Bricostore Via Nettunanse Ariccia 00072 Ita Carta N° *****341 Nessuna Commissione;E;41.58
9;Pag. Del 08/04/23 Ora 23:11 Presso: Sam Beer Viale Dell'aeronautica, Roma 00144 Ita Carta N° *****341 Nessuna Commissione;E;5.00
`

	if expected != sampleOut {
		t.Error("Unexpected result")
	}
}
