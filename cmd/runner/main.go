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
		logrus.Fatalf("configuration loading failed, %s", err.Error())
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// 初始化WorkerPool
	logrus.WithContext(ctx).Infof("initialize worker pool %d", config.WorkerPoolSize)
	worker.InitWorkerPool(config.WorkerPoolSize)

	app, err := runner.NewRunner(config)
	if err != nil {
		logrus.Fatalf("runner creation failed, %s", err.Error())
	}

	// 注册
	if err := app.Register(ctx); err != nil {
		logrus.Fatalf("runner registration failed, %s", err.Error())
	}

	logrus.WithContext(ctx).Infof("runner runing...")
	go app.Run(ctx)

	<-ctx.Done()

	// 等待所有worker执行完成
	logrus.WithContext(ctx).Infof("wait for all workers to complete")
	worker.Pool.WithDone(10)

	logrus.WithContext(ctx).Infof("runner stop")
}
