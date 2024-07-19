package web_server

import (
	"context"
	v1 "github.com/hl540/malou/proto/v1"
)

func (w *WebServer) RunnerInfo(context.Context, *v1.RunnerInfoReq) (*v1.RunnerListResp, error) {
	return nil, nil
}
