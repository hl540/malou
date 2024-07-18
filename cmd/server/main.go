package main

import (
	"context"
	"encoding/json"
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

var runnerInfo = make(map[string]*v1.HeartbeatReq)

func (s *server) Heartbeat(ctx context.Context, req *v1.HeartbeatReq) (*v1.HeartbeatResp, error) {
	runnerInfo[req.Token] = req
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
	http.HandleFunc("/info", func(writer http.ResponseWriter, request *http.Request) {
		jsonByte, _ := json.Marshal(runnerInfo)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(jsonByte)
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
