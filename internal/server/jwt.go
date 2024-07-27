package server

import "github.com/golang-jwt/jwt/v5"

type RunnerRegisterClaims struct {
	jwt.RegisteredClaims
	RunnerID   int64  `json:"runner_id"`
	RunnerCode string `json:"runner_code"`
}
