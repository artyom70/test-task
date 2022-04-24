package price

import (
	"context"
	// "task-artyom/price-backend/service"
)

type Service interface {
	GetPriceList(ctx context.Context, req GetPriceListRequest) (GetPriceListResponse, error)
}

type GetPriceListRequest struct {
	FSYMS []string `json:"fsyms"`
	TSYMS []string `json:"tsyms"`
}

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
	Data   map[string]PriceData `json:"data,omitempty"`
	ErrMsg string               `json:"err,omitempty"`
}
