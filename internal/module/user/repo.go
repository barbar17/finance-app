package user

import (
	"context"
	"database/sql"
	"errors"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) FindByUsername(ctx context.Context, username string) (*User, error) {
	const query = `SELECT id, password, created_at FROM users WHERE username = ?`
	var user User

	user.Username = username
	err := r.db.QueryRowContext(ctx, query, username).Scan(&user.ID, &user.Password, &user.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}
