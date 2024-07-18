package main

import (
	"context"
	"github.com/hl540/malou/internal/runner"
	"github.com/hl540/malou/internal/runner/worker"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os/signal"
	"syscall"
)

var Logger = logrus.New()

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	//加载环境变量
	err := godotenv.Load()
	if err != nil {
		Logger.WithContext(ctx).Warning("Didn't try and open .env by default")
	}

	// 加载配置
	config, err := runner.LoadConfig()
	if err != nil {
		panic(err)
	}

	// 初始化WorkerPool
	Logger.WithContext(ctx).Infof("initialize worker pool %d", config.WorkerPoolSize)
	worker.InitWorkerPool(config.WorkerPoolSize)

	app, err := runner.NewRunner(config)
	if err != nil {
		panic(err)
	}

	Logger.WithContext(ctx).Infof("runner runing...")
	go app.Run(ctx)

	<-ctx.Done()

	Logger.WithContext(ctx).Infof("runner stop")
}
