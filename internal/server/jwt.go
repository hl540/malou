package server

import "github.com/golang-jwt/jwt/v5"

type RunnerRegisterClaims struct {
	jwt.RegisteredClaims
	RunnerID string `json:"runner_id"`
}
