package service

import (
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/repository"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/service/configuration"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/service/image"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/service/orders"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/service/user"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/service/user_information"
)

type Provider struct {
	cfg                    *Config
	repository             repository.Repository
	userService            *user.Service
	userInformationService *user_information.Service
	imageService           *image.Service
	UserConfiguration      *configuration.Service
	OrderService           *orders.Service
}

func NewProvider(cfg *Config, repo repository.Repository) (*Provider, error) {
	userService, err := user.NewService(repo)
	if err != nil {
		return nil, err
	}
	userinformationService, err := user_information.NewService(repo)
	if err != nil {
		return nil, err
	}
	imageService, err := image.NewService(repo)
	if err != nil {
		return nil, err
	}
	userConfigurationService, err := configuration.NewService(repo)
	if err != nil {
		return nil, err
	}
	ordersService, err := orders.NewService(repo)
	if err != nil {
		return nil, err
	}
	return &Provider{
		cfg:                    cfg,
		repository:             repo,
		userService:            userService,
		userInformationService: userinformationService,
		imageService:           imageService,
		UserConfiguration:      userConfigurationService,
		OrderService:           ordersService,
	}, nil
}

func (p *Provider) GetUserService() *user.Service {
	return p.userService
}

func (p *Provider) GetUserInformationService() *user_information.Service {
	return p.userInformationService
}

func (p *Provider) GetImageService() *image.Service {
	return p.imageService
}
func (p *Provider) GetUserConfigurationService() *configuration.Service {
	return p.UserConfiguration
}

func (p *Provider) GetOrdersService() *orders.Service {
	return p.OrderService
}
func (p *Provider) GetConfig() *Config {
	return p.cfg
}

func (p *Provider) Shutdown() {
	p.repository.Shutdown()
}
