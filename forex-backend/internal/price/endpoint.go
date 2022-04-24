package price

import (
	"context"

	"forex-backend/internal"
)

func GetPriceEndpoint(svc Service) internal.EndpointFunc {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(GetPriceListRequest)

		res, err := svc.GetPriceList(ctx, req)
		if err != nil {
			return nil, err
		}

		return res, nil
	}

}
