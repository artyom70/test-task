package service

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

func (r *repository) GetPriceList(ctx context.Context, req GetPriceListRequest) ([]PriceRecord, error) {
	sql := `SELECT from_currency, to_currency, data FROM asset WHERE from_currency IN(?) AND to_currency IN(?)`

	query, args, err := sqlx.In(sql, req.FSYMS, req.TSYMS)
	if err != nil {
		return nil, err
	}

	query = r.db.Rebind(query)
	rows, err := r.db.QueryxContext(ctx, query, args...)

	defer rows.Close()

	var res []PriceRecord

	for rows.Next() {
		var record PriceRecord
		if err := rows.StructScan(&record); err != nil {
			return nil, err
		}

		res = append(res, record)
	}

	return res, nil

}
