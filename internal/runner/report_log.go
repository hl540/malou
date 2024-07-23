package runner

import (
	"fmt"
	v1 "github.com/hl540/malou/proto/v1"
	"github.com/sirupsen/logrus"
	"time"
)

type ReportLog struct {
	pipelineID   string
	reportStream v1.Malou_ReportPipelineLogClient
	step         string
	cmd          string
	timestamp    int64
}

func NewReportLog(pipelineID string, reportStream v1.Malou_ReportPipelineLogClient) *ReportLog {
	return &ReportLog{
		pipelineID:   pipelineID,
		reportStream: reportStream,
		timestamp:    time.Now().Unix(),
	}
}

func (l *ReportLog) WithStep(name string) *ReportLog {
	return &ReportLog{
		pipelineID:   l.pipelineID,
		reportStream: l.reportStream,
		step:         name,
		timestamp:    l.timestamp,
	}
}

func (l *ReportLog) WithCmd(cmd string) *ReportLog {
	return &ReportLog{
		pipelineID:   l.pipelineID,
		reportStream: l.reportStream,
		step:         l.step,
		cmd:          cmd,
		timestamp:    l.timestamp,
	}
}

func (l *ReportLog) Send(req *v1.PipelineLog) {
	req.PipelineId = l.pipelineID
	req.Step = l.step
	req.Cmd = l.cmd
	req.Timestamp = time.Now().Unix()
	req.Duration = req.Timestamp - l.timestamp
	if l.reportStream != nil {
		if err := l.reportStream.Send(req); err != nil {
			logrus.Errorf("Failed to report log, %s", err.Error())
		}
	}
}

func (l *ReportLog) Log(message string, v ...any) {
	l.Send(&v1.PipelineLog{
		Type:    v1.PipelineLogType_LOG,
		Message: fmt.Sprintf(message, v...),
	})
}

func (l *ReportLog) Error(message string, v ...any) {
	l.Send(&v1.PipelineLog{
		Type:    v1.PipelineLogType_ERROR,
		Message: fmt.Sprintf(message, v...),
	})
}

func (l *ReportLog) Done(message string, v ...any) {
	l.Send(&v1.PipelineLog{
		Type:    v1.PipelineLogType_DONE,
		Message: fmt.Sprintf(message, v...),
	})
}
