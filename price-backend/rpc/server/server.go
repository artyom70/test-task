package server

import (
	"log"

	"github.com/nats-io/nats.go"
)

type server struct {
	susbcription nats.Subscription
	routes       map[string]nats.MsgHandler
}

// NewServer creates instance of rpc server based on routes
func NewServer(routes map[string]nats.MsgHandler) *server {
	return &server{
		routes: routes,
	}
}

// Subscribe subscribes on all routes
func (srv *server) Subscribe(conn nats.Conn, serviceName string) error {
	subs := map[string]*nats.Subscription{}
	var err error
	for subject, handler := range srv.routes {
		sub, err1 := conn.QueueSubscribe(subject, serviceName, handler)
		if err1 != nil {
			err = err1
			break
		}

		subs[subject] = sub
	}

	if err != nil {
		for subj, s := range subs {
			if err := s.Unsubscribe(); err != nil {
				log.Printf("could not unsubscribe from %s: %v", subj, err)
			}
		}
		return err
	}

	return err
}
