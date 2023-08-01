package main

import (
	httpcoding "cart_mircoservice/iternal/coding/http"
	"cart_mircoservice/iternal/config"
	db "cart_mircoservice/iternal/db/redis"
	"cart_mircoservice/iternal/endpoint"
	"cart_mircoservice/iternal/service"
	"errors"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	ErrBadRouting = errors.New("bad routing")
)

func NewService(
	svcEndpoints endpoint.Endpoints, logger log.Logger,
) http.Handler {
	// set-up router and initialize http endpoints
	r := mux.NewRouter()
	options := []kithttp.ServerOption{
		//kithttp.ServerErrorHandler(log.L),
		//kithttp.ServerErrorEncoder(kithttp.ErrorEncoder()),
	}
	// HTTP Post - /orders
	r.Methods("GET").Path("/items/").Handler(kithttp.NewServer(
		svcEndpoints.GetFromCart,
		httpcoding.DecodeGetFromCart,
		httpcoding.EncodeResponse,
		options...,
	))

	// HTTP Post - /orders/{id}
	r.Methods("POST").Path("/items/").Handler(kithttp.NewServer(
		svcEndpoints.AddToCart,
		httpcoding.DecodeAddToCartRequest,
		httpcoding.EncodeResponse,
		options...,
	))

	// HTTP Post - /orders/status
	r.Methods("DELETE").Path("/items/").Handler(kithttp.NewServer(
		svcEndpoints.DeleteFromCart,
		httpcoding.DecodeDeleteFromCartRequest,
		httpcoding.EncodeResponse,
		options...,
	))
	return r
}
func main() {
	var (
		httpAddr = flag.String("http.addr", ":8000", "HTTP listen address")
	)
	flag.Parse()
	cfg := config.MustLoad()
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "cart",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	clientRedis := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Address,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Db,
	})

	redis := db.New(clientRedis, logger)
	svc := service.New(redis, logger)
	var h http.Handler
	{
		endpoints := endpoint.Make(svc)
		h = NewService(endpoints, logger)
	}
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	go func() {
		level.Info(logger).Log("transport", "HTTP", "addr", *httpAddr)
		server := &http.Server{
			Addr:    *httpAddr,
			Handler: h,
		}
		errs <- server.ListenAndServe()
	}()
	level.Error(logger).Log("exit", <-errs)

}
