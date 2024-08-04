package runner_server

import (
	"fmt"
	"github.com/hl540/malou/internal/server/storage"
	v1 "github.com/hl540/malou/proto/v1"
	"github.com/sirupsen/logrus"
	"io"
)

func (s *RunnerServer) ReportPipelineLog(stream v1.Malou_ReportPipelineLogServer) error {
	var serial int64
	for {
		serial++
		reportLog, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&v1.ReportPipelineLogResp{})
		}
		if err != nil {
			logrus.WithContext(stream.Context()).Errorf(err.Error())
			return err
		}
		fmt.Printf("[ReportPipelineLog] %v\n", reportLog)
		err = storage.PipelineInstanceLog.Insert(stream.Context(), &storage.PipelineInstanceLogModel{
			PipelineInstanceId: reportLog.PipelineInstanceId,
			Serial:             serial,
			StepName:           reportLog.StepName,
			Cmd:                reportLog.Cmd,
			Result:             reportLog.Result,
			Type:               reportLog.Type.String(),
			Timestamp:          reportLog.Timestamp,
			Duration:           reportLog.Duration,
		})
		if err != nil {
			logrus.WithContext(stream.Context()).Warningf("ReportPipelineLog err: %s", err.Error())
		}
	}
}
