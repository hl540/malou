package runner_server

import (
	"fmt"
	v1 "github.com/hl540/malou/proto/v1"
	"io"
	"time"
)

func (s *RunnerServer) ReportPipelineLog(stream v1.Malou_ReportPipelineLogServer) error {
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
