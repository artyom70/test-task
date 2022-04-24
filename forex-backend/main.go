package main

import (
	"context"
	"fmt"
	"forex-backend/internal/price"
	"forex-backend/internal/price/rpc/client"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/nats-io/nats.go"

	"github.com/oklog/run"
)

var (
	argHTTPPort = mustEnv("HTTP_PORT")
	argNatsHost = mustEnv("NATS_HOST")
	argNatsPort = mustEnv("NATS_PORT")
)

func main() {
	var g run.Group

	conn := connectNats(argNatsHost, argNatsPort)
	svc := client.New(conn)
	handler := price.MakeHandler(svc)

	router := chi.NewRouter()
	router.Mount("/", handler)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", argHTTPPort),
		Handler: router,
	}

	g.Add(func() error {
		log.Printf("[http_server] started listening on port %s \n", argHTTPPort)
		return srv.ListenAndServe()
	}, func(err error) {
		srv.Shutdown(context.Background())
	})

	g.Add(func() error {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		c := <-sigChan
		return fmt.Errorf("terminated with sig %q", c)
	}, func(err error) {
		//
	})

	fmt.Println("price-backend terminated with error:", g.Run())

}

func mustEnv(name string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		panic(fmt.Sprintf("%s is required", name))
	}

	return value
}

func connectNats(host string, port string) *nats.Conn {
	url := fmt.Sprintf("%s:%s", host, port)
	nc, err := nats.Connect(url)
	if err != nil {
		panic(fmt.Sprintf("couldn't connect to nats, error: %v, url: %v", err, url))
	}

	return nc
}
