package web_server

import (
	"context"
	"github.com/hl540/malou/internal/server/storage"
	v1 "github.com/hl540/malou/proto/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (w *WebServer) UpdateRunner(ctx context.Context, req *v1.UpdateRunnerReq) (*v1.UpdateRunnerResp, error) {
	newRunner := &storage.RunnerModel{
		ID:   req.Id,
		Code: req.Code,
		Name: req.Name,
	}
	runner, err := storage.Runner.GetByID(ctx, newRunner.ID)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	err = storage.TransactionCtx(ctx, func(ctx context.Context, tx storage.Session) error {
		runnerDao := storage.NewRunnerDao(tx)
		if err := runnerDao.Update(ctx, newRunner); err != nil {
			return err
		}
		if err := runnerDao.SaveLabel(ctx, runner.ID, req.Labels); err != nil {
			return err
		}
		if err := runnerDao.SaveEnv(ctx, runner.ID, req.Env); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &v1.UpdateRunnerResp{Id: runner.ID, Code: runner.Code}, nil
}
