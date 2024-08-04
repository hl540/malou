package runner

import (
	"fmt"
	v1 "github.com/hl540/malou/proto/v1"
	"github.com/sirupsen/logrus"
	"time"
)

type ReportLog struct {
	pipelineId   string
	reportStream v1.Malou_ReportPipelineLogClient
	step         string
	cmd          string
	timestamp    int64
}

func NewReportLog(pipelineId string, reportStream v1.Malou_ReportPipelineLogClient) *ReportLog {
	return &ReportLog{
		pipelineId:   pipelineId,
		reportStream: reportStream,
		timestamp:    time.Now().Unix(),
	}
}

func (l *ReportLog) WithStep(name string) *ReportLog {
	return &ReportLog{
		pipelineId:   l.pipelineId,
		reportStream: l.reportStream,
		step:         name,
		timestamp:    l.timestamp,
	}
}

func (l *ReportLog) WithCmd(cmd string) *ReportLog {
	return &ReportLog{
		pipelineId:   l.pipelineId,
		reportStream: l.reportStream,
		step:         l.step,
		cmd:          cmd,
		timestamp:    l.timestamp,
	}
}

func (l *ReportLog) Send(req *v1.PipelineInstanceLog) {
	req.PipelineInstanceId = l.pipelineId
	req.StepName = l.step
	req.Cmd = l.cmd
	req.Timestamp = time.Now().Unix()
	req.Duration = req.Timestamp - l.timestamp
	fmt.Printf("[ReportLog] %v\n", req)
	if l.reportStream != nil {
		if err := l.reportStream.Send(req); err != nil {
			logrus.Errorf("Failed to report log, %s", err.Error())
		}
	}
}

func (l *ReportLog) Log(message string, v ...any) {
	l.Send(&v1.PipelineInstanceLog{
		Type:   v1.PipelineLogType_LOG,
		Result: fmt.Sprintf(message, v...),
	})
}

func (l *ReportLog) Error(message string, v ...any) {
	l.Send(&v1.PipelineInstanceLog{
		Type:   v1.PipelineLogType_ERROR,
		Result: fmt.Sprintf(message, v...),
	})
}

func (l *ReportLog) Done(message string, v ...any) {
	l.Send(&v1.PipelineInstanceLog{
		Type:   v1.PipelineLogType_DONE,
		Result: fmt.Sprintf(message, v...),
	})
}
