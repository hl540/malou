package main

import (
	"context"
	"errors"
	"github.com/hl540/malou/proto/v1"
	"github.com/hl540/malou/utils"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

type server struct {
	v1.UnimplementedMalouServerServer
}

func (s *server) Heartbeat(ctx context.Context, req *v1.HeartbeatReq) (*v1.HeartbeatResp, error) {
	log.Printf("[%s]心跳请求", ctx.Value("token"))
	//for i, item := range req.CpuPercent {
	//	log.Printf("[cpu%d]使用率: %.2f", i, item)
	//}
	//log.Printf("内存总量: %v MB", req.MemoryInfo.Total/1024/1024)
	//log.Printf("内存使用量: %v MB", req.MemoryInfo.Used/1024/1024)
	//log.Printf("内存空闲量: %v MB", req.MemoryInfo.Free/1024/1024)
	//log.Printf("内存使用率: %.2f%%", req.MemoryInfo.UsedPercent)
	//
	//log.Printf("磁盘总量: %v MB", req.DiskInfo.Total/1024/1024)
	//log.Printf("磁盘使用量: %v MB", req.DiskInfo.Used/1024/1024)
	//log.Printf("磁盘空闲量: %v MB", req.DiskInfo.Free/1024/1024)
	//log.Printf("磁盘使用率: %.2f%%", req.DiskInfo.UsedPercent)
	return &v1.HeartbeatResp{
		Timestamp: time.Now().Unix(),
		Message:   "success",
	}, nil
}

var pipeline = make(chan *v1.Pipeline)

func (s *server) PullPipeline(context.Context, *v1.PullPipelineReq) (*v1.PullPipelineResp, error) {
	select {
	case p := <-pipeline:
		return &v1.PullPipelineResp{
			PipelineId: utils.StringWithCharset(20, utils.Charset2),
			Pipeline:   p,
		}, nil
	default:
		return nil, errors.New("not")
	}
}

func (s *server) ReportPipelineLog(stream v1.MalouServer_ReportPipelineLogServer) error {
	for {
		reportLog, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&v1.ReportPipelineLogResp{
				Timestamp: time.Now().Unix(),
				Message:   "success",
			})
		}
		if err != nil {
			return err
		}
		log.Println(reportLog)
	}
}

func main() {
	http.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
		pipeline <- &v1.Pipeline{
			Kind: "docker",
			Type: "xxxx",
			Name: "asada",
			Steps: []*v1.Step{
				{
					Name:  "checkout",
					Image: "alpine:3.18",
					Commands: []string{
						"echo $(uname -a)",
						"echo $(pwd)",
						"echo 123546 > 123456.txt",
					},
				},
				{
					Name:  "build",
					Image: "alpine:3.18",
					Commands: []string{
						"echo abc > abc.txt",
						"echo def > def.txt",
						"ls -l -a",
						"sleep 10",
					},
				},
			},
		}
	})
	go http.ListenAndServe(":9999", nil)

	lis, err := net.Listen("tcp", ":5555")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	v1.RegisterMalouServerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
