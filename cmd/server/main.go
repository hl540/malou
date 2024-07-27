package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hl540/malou/internal/server"
	"github.com/hl540/malou/internal/server/runner_server"
	"github.com/hl540/malou/internal/server/storage"
	"github.com/hl540/malou/internal/server/web_server"
	"github.com/hl540/malou/proto/v1"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"os/signal"
	"syscall"
)

func main() {
	// 加载配置
	config, err := server.LoadConfig()
	if err != nil {
		panic(err)
	}

	// 初始化数据库
	db, err := storage.InitDB(config)
	if err != nil {
		panic(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	grpcAddress := fmt.Sprintf("%s:%d", config.GrpcHost, config.GrpcPort)
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	v1.RegisterMalouServer(s, &runner_server.RunnerServer{})
	v1.RegisterMalouWebServer(s, &web_server.WebServer{})

	logrus.WithContext(ctx).Infof("Serving gRPC on %s", grpcAddress)
	go func() {
		if err := s.Serve(lis); err != nil {
			logrus.WithContext(ctx).Errorf("Failed to listen %s", err.Error())
		}
	}()

	// 创建grpc client
	conn, err := grpc.NewClient(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	mux := runtime.NewServeMux()
	if err := v1.RegisterMalouWebHandler(ctx, mux, conn); err != nil {
		panic(err)
	}

	httpAddress := fmt.Sprintf("%s:%d", config.HttpHost, config.HttpPort)
	gwServer := &http.Server{
		Addr:    httpAddress,
		Handler: mux,
	}
	logrus.WithContext(ctx).Infof("Serving gRPC-Gateway on http://%s", httpAddress)
	go func() {
		if err := gwServer.ListenAndServe(); err != nil {
			logrus.WithContext(ctx).Errorf("Failed to listen %s", err.Error())
		}
	}()

	<-ctx.Done()

	if err := db.Close(); err != nil {
		logrus.WithContext(ctx).Errorf("Failed to close database: %s", err.Error())
	}
	logrus.WithContext(ctx).Infof("RunnerServer stop")
}
