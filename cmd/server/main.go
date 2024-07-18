package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hl540/malou/internal/server"
	"github.com/hl540/malou/proto/v1"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"os/signal"
	"syscall"
)

var Logger = logrus.New()

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	lis, err := net.Listen("tcp", ":5555")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	v1.RegisterMalouServer(s, &server.Server{})
	v1.RegisterMalouWebServer(s, &server.WebServer{})

	Logger.WithContext(ctx).Info("Serving gRPC on 0.0.0.0:5555")
	go func() {
		if err := s.Serve(lis); err != nil {
			Logger.WithContext(ctx).Errorf("Failed to listen %s", err.Error())
		}
	}()

	// 创建grpc client
	conn, err := grpc.NewClient("0.0.0.0:5555", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	mux := runtime.NewServeMux()
	if err := v1.RegisterMalouWebHandler(ctx, mux, conn); err != nil {
		panic(err)
	}

	gwServer := &http.Server{Addr: ":6666", Handler: mux}
	Logger.WithContext(ctx).Info("Serving gRPC-Gateway on http://0.0.0.0:6666")
	go func() {
		if err := gwServer.ListenAndServe(); err != nil {
			Logger.WithContext(ctx).Errorf("Failed to listen %s", err.Error())
		}
	}()

	<-ctx.Done()
	Logger.WithContext(ctx).Infof("Server stop")
}
