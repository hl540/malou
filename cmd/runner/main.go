package main

import (
	"context"
	"github.com/hl540/malou/internal/runner"
	"github.com/hl540/malou/internal/runner/worker"
	"github.com/sirupsen/logrus"
	"os/signal"
	"syscall"
)

func main() {
	// 加载配置
	config, err := runner.LoadConfig()
	if err != nil {
		panic(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// 初始化WorkerPool
	logrus.WithContext(ctx).Infof("initialize worker pool %d", config.WorkerPoolSize)
	worker.InitWorkerPool(config.WorkerPoolSize)

	app, err := runner.NewRunner(config)
	if err != nil {
		panic(err)
	}

	logrus.WithContext(ctx).Infof("runner runing...")
	go app.Run(ctx)

	<-ctx.Done()

	logrus.WithContext(ctx).Infof("runner stop")
}
