package api

import "context"

type Client interface {
	GetPrice(ctx context.Context, fsyms []string, tsyms []string) (PriceResponse, error)
}

type PriceResponse struct {
	Data map[string]map[string]PriceData `json:"DISPLAY"`
}

type PriceData struct {
	Change24Hour    string `json:"CHANGE24HOUR"`
	ChangePCT24Hour string `json:"CHANGEPCT24HOUR"`
	Open24Hour      string `json:"OPEN24HOUR"`
	Volume24Hour    string `json:"VOLUME24HOUR"`
	Volume24HourTo  string `json:"VOLUME24HOURTO"`
	Low24Hour       string `json:"LOW24HOUR"`
	High24Hour      string `json:"HIGH24HOUR"`
	Price           string `json:"PRICE"`
	Supply          string `json:"SUPPLY"`
	MKTCap          string `json:"MKTCAP"`
}
