package user

import "time"

type User struct {
	ID        int64
	Username  string
	Password  string
	Name      string
	CreatedAt time.Time
}
