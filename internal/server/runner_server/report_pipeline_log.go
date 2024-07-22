package runner_server

import (
	"github.com/hl540/malou/internal/server/storage"
	v1 "github.com/hl540/malou/proto/v1"
	"github.com/sirupsen/logrus"
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
			logrus.WithContext(stream.Context()).Errorf(err.Error())
			return err
		}
		err = storage.AddPipelineLog(stream.Context(), &storage.PipelineLog{
			PipelineID: reportLog.PipelineId,
			Step:       reportLog.Step,
			Cmd:        reportLog.Cmd,
			Message:    reportLog.Message,
			Type:       reportLog.Type.String(),
			Timestamp:  reportLog.Timestamp,
			Duration:   reportLog.Duration,
		})
		if err != nil {
			logrus.WithContext(stream.Context()).Errorf("[AddPipelineLog] err: %s", err.Error())
		}
	}
}
