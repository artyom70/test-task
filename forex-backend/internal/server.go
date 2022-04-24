package internal

import (
	"context"
	"net/http"
)

type DecodeRequestFunc func(ctx context.Context, r *http.Request) (interface{}, error)
type EncodeRequestFunc func(ctx context.Context, rw http.ResponseWriter, res interface{}) error
type EndpointFunc func(ctx context.Context, rw interface{}) (interface{}, error)

type server struct {
	e   EndpointFunc
	dec DecodeRequestFunc
	enc EncodeRequestFunc
}

func NewServer(e EndpointFunc, dec DecodeRequestFunc, enc EncodeRequestFunc) *server {
	return &server{
		e:   e,
		dec: dec,
		enc: enc,
	}
}

func (s *server) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	req, err := s.dec(r.Context(), r)
	if err != nil {
		EncodeError(err, rw)
		return
	}

	res, err := s.e(r.Context(), req)
	if err != nil {
		EncodeError(err, rw)
		return
	}

	if err := s.enc(r.Context(), rw, res); err != nil {
		EncodeError(err, rw)
		return
	}

}
