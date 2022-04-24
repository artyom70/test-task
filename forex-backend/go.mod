module forex-backend

go 1.16

require (
	github.com/go-chi/chi/v5 v5.0.7
	github.com/nats-io/nats.go v1.14.0
	github.com/oklog/run v1.1.0
)

replace github.com/price-backend => ../price-backend
