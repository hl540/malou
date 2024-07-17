package main

import (
	"context"
	"github.com/hl540/malou/internal/app/agent"
	"github.com/hl540/malou/internal/worker"
	"github.com/joho/godotenv"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	//加载环境变量
	err := godotenv.Load()
	if err != nil {
		log.Println("Didn't try and open .env by default")
	}

	// 加载配置
	config, err := agent.LoadConfig()
	if err != nil {
		panic(err)
	}

	// 初始化WorkerPool
	worker.InitWorkerPool(config.WorkerPoolSize)

	app, err := agent.NewAgent(config)
	if err != nil {
		panic(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go app.Run(ctx)

	agent.Logger.WithContext(ctx).Infof("agent runing...")

	<-ctx.Done()

	agent.Logger.WithContext(ctx).Infof("agent stop")

	//pipeline := &agent.Pipeline{
	//	Kind: "docker",
	//	Type: "xxxx",
	//	Name: "asada",
	//	Steps: []*agent.Step{
	//		{
	//			Name:  "checkout",
	//			Image: "alpine:3.18",
	//			Commands: []string{
	//				"echo $(uname -a)",
	//				"echo $(pwd)",
	//				"echo 123546 > 123456.txt",
	//			},
	//		},
	//		{
	//			Name:  "build",
	//			Image: "alpine:3.18",
	//			Commands: []string{
	//				"echo abc > abc.txt",
	//				"echo def > def.txt",
	//				"ls -l -a",
	//			},
	//		},
	//	},
	//}
	//agent.ExecutePipeline(pipeline)
}
