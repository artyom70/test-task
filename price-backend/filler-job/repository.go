package fillerjob

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sql.DB) *repository {
	dbx := sqlx.NewDb(db, "postgres")
	return &repository{
		db: dbx,
	}
}

func (r *repository) GetCurrencyPairs(ctx context.Context, fromCurrency []string, toCurrency []string) ([]CurrencyPairs, error) {
	sql := `SELECT from_currency, to_currency FROM asset WHERE from_currency IN(?) AND to_currency IN(?) AND is_acquired = false`

	query, args, err := sqlx.In(sql, fromCurrency, toCurrency)
	if err != nil {
		return nil, err
	}

	query = r.db.Rebind(query)
	rows, err := r.db.QueryxContext(ctx, query, args...)

	defer rows.Close()

	var res []CurrencyPairs

	for rows.Next() {
		var record CurrencyPairs
		if err := rows.StructScan(&record); err != nil {
			return nil, err
		}

		res = append(res, record)
	}

	return res, nil
}
