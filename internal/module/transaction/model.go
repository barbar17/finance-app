package transaction

import "time"

type Transaction struct {
	ID        int
	Name      string
	Amount    int
	Desc      string
	CreatedAt time.Time
}
