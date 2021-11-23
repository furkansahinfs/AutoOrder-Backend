package service

import "github.com/furkansahinfs/AutoOrder-Backend/pkg/service/user"

type Config struct{}

type Service interface {
	GetConfig() *Config
	GetUserService() *user.Service
	Shutdown()
}
