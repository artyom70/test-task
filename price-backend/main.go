package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"price-backend/rpc/server"
	"price-backend/service"
	"price-backend/service/api"
	"syscall"

	"github.com/oklog/run"

	nats "github.com/nats-io/nats.go"
)

var (
	argNatsHost    = mustEnv("NATS_HOST")
	argNatsPort    = mustEnv("NATS_PORT")
	argServiceName = mustEnv("SERVICE_NAME")
	argAPIBaseURL  = mustEnv("API_BASE_URL")
)

func main() {
	var g run.Group
	log.SetFlags(log.LstdFlags)

	conn := connectNats(argNatsHost, argNatsPort)
	apiClient := api.NewAPI(http.DefaultClient, argAPIBaseURL)
	repo := service.NewRepository(nil)
	svc := service.New(apiClient, repo)

	priceListener := server.NewListener(conn, svc, argServiceName)

	g.Add(func() error {
		log.Printf("[price-rpc] started listening")
		return priceListener.ListenAndServe()
	}, func(err error) {
		priceListener.ShutDown(context.Background())
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
