package transaction

import (
	"context"
	"database/sql"

	"github.com/barbar17/finance-app/internal/types"
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
		"created_at": "created_at",
	}

	sortColumn, ok := allowedSort[p.Sort]
	if !ok {
		sortColumn = "created_at"
	}

	if p.Order != "asc" && p.Order != desc {
		p.Order = "desc"
	}

	searchPattern := "%" + p.Search + "%"
}
