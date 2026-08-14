package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/barbar17/finance-app/internal/types"
	"github.com/barbar17/finance-app/internal/utils"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) GetUsersTable(ctx context.Context, p types.TableParams) ([]User, int, error) {
	allowedSort := map[string]string{
		"username":   "username",
		"name":       "name",
		"created_at": "created_at",
	}
	sortColumn, orderColumn, searchPattern := utils.TableParamsInit(allowedSort, p.Sort, p.Order, p.Search)

	countQuery := `SELECT COUNT(*) FROM users WHERE username LIKE ? OR Name LIKE ?`
	var total int
	if err := r.db.QueryRowContext(ctx, countQuery, searchPattern, searchPattern).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := fmt.Sprintf(`
		SELECT id, username, name, created_at FROM users
		WHERE username LIKE ? OR name LIKE ?
		ORDER BY %s %s LIMIT ? OFFSET ?
	`, sortColumn, orderColumn)

	rows, err := r.db.QueryContext(ctx, query, searchPattern, searchPattern, p.Limit, p.Offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Username, &user.Name, &user.CreatedAt)
		if err != nil {
			return nil, 0, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *Repo) FindByUsername(ctx context.Context, username string) (*User, error) {
	const query = `SELECT id, password, name, created_at FROM users WHERE username = ?`
	var user User

	user.Username = username
	err := r.db.QueryRowContext(ctx, query, username).Scan(&user.ID, &user.Password, &user.Name, &user.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}
