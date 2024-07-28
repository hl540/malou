package test

import (
	"context"
	"github.com/google/uuid"
	"github.com/hl540/malou/internal/server/storage"
	"testing"
)

func TestPipelineLog(t *testing.T) {
	err := storage.AddPipelineLog(context.Background(), &storage.PipelineLog{
		PipelineID: uuid.New().String(),
		Step:       "build",
		Cmd:        "go build -o ./server main.go",
		Message:    "run go build",
		Timestamp:  1721615868,
		Duration:   10,
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("success")
}

func TestGetPipelineLogs(t *testing.T) {
	logs, err := storage.GetPipelineLogs(context.Background(), "9BEV1L96RSL0CFHM29KX", 11)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(logs)
}
