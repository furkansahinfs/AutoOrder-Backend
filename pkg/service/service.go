package service

import (
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/service/user"
	userinformation "github.com/furkansahinfs/AutoOrder-Backend/pkg/service/userInformation"
)

type Config struct{}

type Service interface {
	GetConfig() *Config
	GetUserService() *user.Service
	GetUserInformationService() *userinformation.Service
	Shutdown()
}
