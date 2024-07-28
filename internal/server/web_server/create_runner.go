package web_server

import (
	"context"
	"github.com/google/uuid"
	"github.com/hl540/malou/internal/server/storage"
	v1 "github.com/hl540/malou/proto/v1"
	"github.com/hl540/malou/utils"
)

func (w *WebServer) CreateRunner(ctx context.Context, req *v1.CreateRunnerReq) (*v1.CreateRunnerResp, error) {
	runner := &storage.RunnerModel{
		Code:   uuid.NewString(),
		Secret: utils.StringWithCharsetV4(20),
		Name:   req.Name,
	}
	err := storage.TransactionCtx(ctx, func(ctx context.Context, tx storage.Session) error {
		runnerDao := storage.NewRunnerDao(tx)
		if err := runnerDao.Create(ctx, runner); err != nil {
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
		return nil, err
	}
	return &v1.CreateRunnerResp{Id: runner.ID, Code: runner.Code}, nil
}
