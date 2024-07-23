package runner_server

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hl540/malou/internal/server"
	"github.com/hl540/malou/internal/server/storage"
	v1 "github.com/hl540/malou/proto/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (s *RunnerServer) RegisterRunner(ctx context.Context, req *v1.RegisterRunnerReq) (*v1.RegisterRunnerResp, error) {
	runner, err := storage.GetRunnerByID(ctx, req.Token)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	claims := server.RunnerRegisterClaims{
		RunnerID: req.Token,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 1)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtString, err := token.SignedString([]byte(runner.Key))
	if err != nil {
		return nil, err
	}
	return &v1.RegisterRunnerResp{Jwt: jwtString}, nil
}
