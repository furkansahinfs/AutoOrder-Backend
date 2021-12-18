package service

import (
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/service/user"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/service/user_information"
)

type Config struct{}

type Service interface {
	GetConfig() *Config
	GetUserService() *user.Service
	GetUserInformationService() *user_information.Service
	Shutdown()
}
