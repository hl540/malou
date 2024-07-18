package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/hl540/malou/proto/v1"
	"github.com/hl540/malou/utils"
	"github.com/sirupsen/logrus"
	"io"
	"time"
)

var Logger = logrus.New()

type Server struct {
	v1.UnimplementedMalouServer
}

func (s *Server) Heartbeat(ctx context.Context, req *v1.HeartbeatReq) (*v1.HeartbeatResp, error) {
	//Logger.WithContext(ctx).Infof("[Heartbeat] %v", req)
	return &v1.HeartbeatResp{Timestamp: time.Now().Unix(), Message: "Received"}, nil
}

var pipeline = &v1.Pipeline{
	Kind: "docker",
	Type: "xxxx",
	Name: utils.StringWithCharsetV3(8),
	Steps: []*v1.Step{
		{
			Name:  "checkout",
			Image: "alpine:3.18",
			Commands: []string{
				"ls -l -a",
				"echo $(pwd)",
				"echo $(uname -a) > log.txt",
				fmt.Sprintf("echo %s >> %s.txt", utils.StringWithCharsetV4(100), utils.StringWithCharsetV3(10)),
				fmt.Sprintf("echo %s >> %s.txt", utils.StringWithCharsetV4(100), utils.StringWithCharsetV3(10)),
				fmt.Sprintf("echo %s >> %s.txt", utils.StringWithCharsetV4(100), utils.StringWithCharsetV3(10)),
				fmt.Sprintf("echo %s >> %s.txt", utils.StringWithCharsetV4(100), utils.StringWithCharsetV3(10)),
				"echo $(ls -l -a) > log.txt",
				"sleep 10",
			},
		},
		{
			Name:  "build",
			Image: "alpine:3.18",
			Commands: []string{
				"echo $(pwd)",
				"ls -l -a",
			},
		},
	},
}

var t = time.NewTimer(1 * time.Second)

func (s *Server) PullPipeline(ctx context.Context, req *v1.PullPipelineReq) (*v1.PullPipelineResp, error) {
	select {
	case <-t.C:
		Logger.WithContext(ctx).Infof("[PullPipeline] %v", req)
		return &v1.PullPipelineResp{
			PipelineId: utils.StringWithCharsetV2(20),
			Pipeline:   pipeline,
		}, nil
	default:
		return nil, errors.New("not")
	}
}

func (s *Server) ReportPipelineLog(stream v1.Malou_ReportPipelineLogServer) error {
	for {
		reportLog, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&v1.ReportPipelineLogResp{
				Timestamp: time.Now().Unix(),
				Message:   "success",
			})
		}
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		fmt.Println(reportLog.Message)
	}
}
