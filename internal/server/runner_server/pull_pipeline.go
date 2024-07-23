package runner_server

import (
	"context"
	"errors"
	"fmt"
	v1 "github.com/hl540/malou/proto/v1"
	"github.com/hl540/malou/utils"
	"github.com/sirupsen/logrus"
	"time"
)

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

var t = time.NewTicker(5 * time.Second)

func (s *RunnerServer) PullPipeline(ctx context.Context, req *v1.PullPipelineReq) (*v1.PullPipelineResp, error) {
	select {
	case <-t.C:
		logrus.WithContext(ctx).Infof("[PullPipeline] %v", req)
		return &v1.PullPipelineResp{
			PipelineId: utils.StringWithCharsetV2(20),
			Pipeline:   pipeline,
		}, nil
	default:
		return nil, errors.New("not")
	}
}
