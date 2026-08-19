package transaction

import "time"

type Transaction struct {
	ID        int
	Name      string
	Amount    int
	Desc      string
	Category  string
	CreatedAt time.Time
}

type CreateTransaction struct {
	Name     string
	Amount   int
	Desc     string
	Category string
}
