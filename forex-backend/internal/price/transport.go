package price

import (
	"context"
	"fmt"
	"forex-backend/internal"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func MakeHandler(svc Service) chi.Router {
	mux := chi.NewMux()

	priceEndpoint := GetPriceEndpoint(svc)
	priceHandler := internal.NewServer(priceEndpoint, decodePriceRequest, internal.EncodeResponse)
	mux.Method(http.MethodGet, "/price", priceHandler)
	return mux
}

func decodePriceRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	values := r.URL.Query()
	fsyms := values.Get("fsyms")
	tsyms := values.Get("tsyms")
	fsymsArr := []string{}
	tsymsArr := []string{}

	for _, v := range strings.Split(fsyms, ",") {
		v = strings.TrimSpace(v)
		if v != "" {
			fsymsArr = append(fsymsArr, v)
		}
	}

	for _, v := range strings.Split(tsyms, ",") {
		v = strings.TrimSpace(v)
		if v != "" {
			tsymsArr = append(tsymsArr, v)
		}
	}

	var req GetPriceListRequest

	req.FSYMS = fsymsArr
	req.TSYMS = tsymsArr

	if len(req.FSYMS) == 0 {
		return nil, fmt.Errorf("fsyms should be non-empty, %w", internal.ErrBadRequest)
	}

	if len(req.TSYMS) == 0 {
		return nil, fmt.Errorf("tsyms should be non-empty, %w", internal.ErrBadRequest)
	}

	return req, nil
}
