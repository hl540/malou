package test

import (
	"context"
	"github.com/hl540/malou/internal/server/web_server"
	v1 "github.com/hl540/malou/proto/v1"
	"testing"
)

func TestWebServer_CreateRunner(t *testing.T) {
	s := new(web_server.WebServer)
	got, err := s.CreateRunner(context.Background(), &v1.CreateRunnerReq{
		Name:   "测试runner",
		Labels: []string{"label1", "label2"},
		Env:    map[string]string{"env1": "value1", "env2": "value2"},
	})
	if err != nil {
		t.Errorf("CreateRunner() error = %v", err)
		return
	}
	t.Logf("got = %v", got)
}
