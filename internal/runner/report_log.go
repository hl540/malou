package runner

import (
	"fmt"
	v1 "github.com/hl540/malou/proto/v1"
	"time"
)

type ReportLog struct {
	pipelineID   string
	reportStream v1.MalouServer_ReportPipelineLogClient
	step         string
	cmd          string
	timestamp    int64
}

func NewReportLog(pipelineID string, reportStream v1.MalouServer_ReportPipelineLogClient) *ReportLog {
	return &ReportLog{
		pipelineID:   pipelineID,
		reportStream: reportStream,
	}
}

func (l *ReportLog) WithStep(name string) *ReportLog {
	return &ReportLog{
		pipelineID:   l.pipelineID,
		reportStream: l.reportStream,
		step:         name,
	}
}

func (l *ReportLog) WithCmd(cmd string) *ReportLog {
	return &ReportLog{
		pipelineID:   l.pipelineID,
		reportStream: l.reportStream,
		step:         l.step,
		cmd:          cmd,
	}
}

func (l *ReportLog) Send(req *v1.ReportPipelineLogReq) {
	if l.reportStream != nil {
		Logger.Infof("%v", req)
		return
	}
	if req != nil {
		req.PipelineId = l.pipelineID
		req.Step = l.step
		req.Cmd = l.cmd
		req.Timestamp = time.Now().Unix()
		req.Duration = req.Timestamp - l.timestamp
	}
	if l.reportStream != nil {
		if err := l.reportStream.Send(req); err != nil {
			Logger.Errorf("Failed to report log, %s", err.Error())
		}
	}
}

func (l *ReportLog) Log(message string, v ...any) {
	l.Send(&v1.ReportPipelineLogReq{
		Type:    v1.ReportType_LOG,
		Message: fmt.Sprintf(message, v...),
	})
}

func (l *ReportLog) Error(message string, v ...any) {
	l.Send(&v1.ReportPipelineLogReq{
		Type:    v1.ReportType_ERROR,
		Message: fmt.Sprintf(message, v...),
	})
}

func (l *ReportLog) Done(message string, v ...any) {
	l.Send(&v1.ReportPipelineLogReq{
		Type:    v1.ReportType_DONE,
		Message: fmt.Sprintf(message, v...),
	})
}
