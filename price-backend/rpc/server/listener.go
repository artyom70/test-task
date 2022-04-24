package server

import (
	"context"
	"encoding/json"
	"price-backend/service"
	"sync"

	"github.com/nats-io/nats.go"
)

type listener struct {
	nc          *nats.Conn
	serviceName string

	wg sync.WaitGroup

	svc service.Service
}

// NewListener creates instance of rpc listener
func NewListener(nc *nats.Conn, svc service.Service, serviceName string) *listener {
	return &listener{
		nc:          nc,
		serviceName: serviceName,
		svc:         svc,
	}
}

// ListenAndServe starts listening routes
func (l *listener) ListenAndServe() error {
	l.wg.Add(1)
	routes := map[string]nats.MsgHandler{}

	routes["price.GetPriceList"] = func(m *nats.Msg) {
		var req service.GetPriceListRequest

		err := json.Unmarshal(m.Data, &req)
		if err != nil {

		}

		res, err := l.svc.GetPriceList(context.Background(), req)
		if err != nil {
			res.ErrMsg = err.Error()
		}

		resRaw, err := json.Marshal(res)
		if err != nil {
			res.ErrMsg = err.Error()
		}

		_ = m.Respond(resRaw)

	}

	server := NewServer(routes)

	err := server.Subscribe(*l.nc, l.serviceName)
	if err != nil {
		return err
	}

	l.wg.Wait()

	return nil
}

// ShutDown ...
func (l *listener) ShutDown(ctx context.Context) {
	l.wg.Done()
	l.nc.Drain()
}
