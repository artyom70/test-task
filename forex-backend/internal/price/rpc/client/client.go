package client

import (
	"context"
	"encoding/json"
	"fmt"
	"forex-backend/internal/price"

	"github.com/nats-io/nats.go"
)

type client struct {
	conn *nats.Conn
}

func New(conn *nats.Conn) *client {
	return &client{conn: conn}
}

func (c *client) GetPriceList(ctx context.Context, req price.GetPriceListRequest) (price.GetPriceListResponse, error) {
	reqRaw, err := json.Marshal(req)
	if err != nil {
		return price.GetPriceListResponse{}, err
	}

	resp, err := c.conn.RequestWithContext(ctx, "price.GetPriceList", reqRaw)
	if err != nil {
		return price.GetPriceListResponse{}, err
	}

	var res price.GetPriceListResponse

	err = json.Unmarshal(resp.Data, &res)
	if err != nil {
		return price.GetPriceListResponse{}, err
	}

	if res.ErrMsg != "" {
		return price.GetPriceListResponse{}, fmt.Errorf(res.ErrMsg)
	}

	return res, nil
}
