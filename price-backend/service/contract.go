package service

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Repository interface {
	GetPriceList(ctx context.Context, req GetPriceListRequest) ([]PriceRecord, error)
}

// Service ...
type Service interface {
	GetPriceList(ctx context.Context, req GetPriceListRequest) (GetPriceListResponse, error)
}

// GetPriceListRequest ...
type GetPriceListRequest struct {
	FSYMS []string `json:"fsyms"`
	TSYMS []string `json:"tsyms"`
}

// PriceData ...
type PriceData struct {
	Change24Hour    string `json:"change24_hour"`
	ChangePCT24Hour string `json:"change_pct24_hour"`
	Open24Hour      string `json:"open24_hour"`
	Volume24Hour    string `json:"volume24_hour"`
	Volume24HourTo  string `json:"volume24_hour_to"`
	Low24Hour       string `json:"low24_hour"`
	High24Hour      string `json:"high24_hour"`
	Price           string `json:"price"`
	Supply          string `json:"supply"`
	MKTCap          string `json:"mkt_cap"`
}

// GetPriceListResponse ...
type GetPriceListResponse struct {
	Data   map[string]PriceData `json:"data"`
	ErrMsg string               `json:"err"`
}

type PriceRecord struct {
	ID           int64  `db:"id"`
	FromCurrency string `db:"from_currency"`
	ToCurrency   string `db:"to_currency"`
	Data         JSON   `db:"data"`
}

type JSON map[string]interface{}

func (a JSON) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *JSON) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
