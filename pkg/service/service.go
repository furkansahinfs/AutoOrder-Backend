package service

import (
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/service/configuration"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/service/image"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/service/orders"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/service/user"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/service/user_information"
)

type Config struct{}

type Service interface {
	GetConfig() *Config
	GetUserService() *user.Service
	GetUserInformationService() *user_information.Service
	GetImageService() *image.Service
	GetUserConfigurationService() *configuration.Service
	GetOrdersService() *orders.Service
	Shutdown()
}
