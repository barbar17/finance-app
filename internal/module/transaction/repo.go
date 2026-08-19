package transaction

import (
	"context"
	"database/sql"
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

func (r *Repo) GetTransactionTable(ctx context.Context, p types.TableParams) ([]Transaction, int, error) {
	allowedSort := map[string]string{
		"name":       "name",
		"amount":     "amount",
		"desc":       "desc",
		"category":   "desc",
		"created_at": "created_at",
	}

	sortColumn, orderColumn, searchPattern := utils.TableParamsInit(allowedSort, p.Sort, p.Order, p.Search)

	countQuery := `SELECT COUNT(*) FROM transactions WHERE name LIKE ? OR category LIKE ? OR amount LIKE ? OR desc LIKE ?`
	var total int
	if err := r.db.QueryRowContext(ctx, countQuery, searchPattern, searchPattern, searchPattern, searchPattern).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := fmt.Sprintf(`
		SELECT id, name, amount, desc, category, created_at FROM transactions
		WHERE name LIKE ? OR category LIKE ? OR amount LIKE ? OR desc LIKE ?
		ORDER BY %s %s LIMIT ? OFFSET ?
	`, sortColumn, orderColumn)

	rows, err := r.db.QueryContext(ctx, query, searchPattern, searchPattern, searchPattern, searchPattern, p.Limit, p.Offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	transactions := make([]Transaction, 0)
	for rows.Next() {
		var transaction Transaction
		if err = rows.Scan(&transaction.ID, &transaction.Name, &transaction.Amount, &transaction.Desc, &transaction.Category, &transaction.CreatedAt); err != nil {
			return nil, 0, err
		}

		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return transactions, total, nil
}

func (r *Repo) Create(ctx context.Context, t CreateTransaction) error {
	query := `INSERT INTO transactions (name, amount, desc, category) VALUES (?,?,?,?)`

	if _, err := r.db.ExecContext(ctx, query, t.Name, t.Amount, t.Desc, t.Category); err != nil {
		return err
	}

	return nil
}
