package main

import (
	"cart_mircoservice/iternal/config"
	db "cart_mircoservice/iternal/db/redis"
	"cart_mircoservice/iternal/endpoint"
	"cart_mircoservice/iternal/service"
	"cart_mircoservice/iternal/transport"
	pb "cart_mircoservice/iternal/transport/pb"
	"errors"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var (
	ErrBadRouting = errors.New("bad routing")
)

func main() {
	//var (
	//	httpAddr = flag.String("http.addr", ":8000", "HTTP listen address")
	//)
	//flag.Parse()
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
	validate := validator.New()

	endpoints := endpoint.Make(svc, validate)
	grpcServer := transport.NewGRPCServer(endpoints, logger)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	grpcListener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	go func() {
		baseServer := grpc.NewServer()
		//level.Info(logger).Log("transport", "HTTP", "addr", *httpAddr)
		pb.RegisterServiceServer(baseServer, grpcServer)
		//server := &http.Server{
		//	Addr:    *httpAddr,
		//	Handler: h,
		//}
		level.Info(logger).Log("msg", "Server started successfully 🚀")
		baseServer.Serve(grpcListener)
	}()
	level.Error(logger).Log("exit", <-errs)

}
