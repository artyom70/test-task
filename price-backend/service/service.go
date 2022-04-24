package service

import (
	"context"
	"fmt"
	"log"
	"price-backend/api"
)

type priceService struct {
	priceAPI  api.Client
	priceRepo Repository
}

// New ...
func New(priceAPI api.Client, priceRepo Repository) *priceService {
	return &priceService{
		priceAPI:  priceAPI,
		priceRepo: priceRepo,
	}
}

// GetPriceList returns price by currency pairs
func (s *priceService) GetPriceList(ctx context.Context, req GetPriceListRequest) (GetPriceListResponse, error) {
	res := GetPriceListResponse{Data: map[string]PriceData{}}

	apiRes, err := s.priceAPI.GetPrice(ctx, req.FSYMS, req.TSYMS)
	if err == nil {
		makeApiCallResponse(&apiRes, &res)
		return res, nil
	}
	if err != nil {
		log.Println(err)
	}

	// I will get from db price data

	// repoRes, err := s.priceRepo.GetPriceList(ctx, req)
	// if err != nil {
	// 	return GetPriceListResponse{}, nil
	// }

	return res, nil
}

func makeApiCallResponse(apiRes *api.PriceResponse, svcRes *GetPriceListResponse) {
	for fromCurrency, data := range apiRes.Data {
		for toCurrency, priceData := range data {
			svcRes.Data[fmt.Sprintf("%s-%s", fromCurrency, toCurrency)] = PriceData{
				Change24Hour:    priceData.Change24Hour,
				ChangePCT24Hour: priceData.ChangePCT24Hour,
				Open24Hour:      priceData.Open24Hour,
				Volume24Hour:    priceData.Volume24Hour,
				Volume24HourTo:  priceData.Volume24HourTo,
				Low24Hour:       priceData.Low24Hour,
				High24Hour:      priceData.High24Hour,
				Price:           priceData.Price,
				Supply:          priceData.Supply,
				MKTCap:          priceData.MKTCap,
			}
		}

	}
}
