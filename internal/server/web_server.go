package server

import (
	"context"
	v1 "github.com/hl540/malou/proto/v1"
)

type WebServer struct {
	v1.UnimplementedMalouWebServer
}

func (w WebServer) RunnerList(ctx context.Context, req *v1.RunnerListReq) (*v1.RunnerListResp, error) {
	//TODO implement me
	panic("implement me")
}
