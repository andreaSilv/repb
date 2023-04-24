package proc

type Transaction struct {
	DayOfTheMonth int32
	Desc          string
	Currency      string
	Amount        float32
}

type TransactionParser interface {
	Parse() []Transaction
}

const (
	EUR string = "E"
	USD        = "D"
)
