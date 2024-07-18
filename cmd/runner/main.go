package main

import (
	"context"
	"github.com/hl540/malou/internal/runner"
	"github.com/hl540/malou/internal/runner/worker"
	"github.com/joho/godotenv"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	//加载环境变量
	err := godotenv.Load()
	if err != nil {
		runner.Logger.WithContext(ctx).Warning("Didn't try and open .env by default")
	}

	// 加载配置
	config, err := runner.LoadConfig()
	if err != nil {
		panic(err)
	}

	// 初始化WorkerPool
	runner.Logger.WithContext(ctx).Infof("initialize worker pool %d", config.WorkerPoolSize)
	worker.InitWorkerPool(config.WorkerPoolSize)

	app, err := runner.NewRunner(config)
	if err != nil {
		panic(err)
	}

	runner.Logger.WithContext(ctx).Infof("runner runing...")
	go app.Run(ctx)

	<-ctx.Done()

	runner.Logger.WithContext(ctx).Infof("runner stop")
}
