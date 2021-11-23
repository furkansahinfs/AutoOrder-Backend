package service

import (
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/repository"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/service/user"
)

type Provider struct {
	cfg         *Config
	repository  repository.Repository
	userService *user.Service
}

func NewProvider(cfg *Config, repo repository.Repository) (*Provider, error) {
	userService, err := user.NewService(repo)
	if err != nil {
		return nil, err
	}
	return &Provider{
		cfg:         cfg,
		repository:  repo,
		userService: userService,
	}, nil
}

func (p *Provider) GetUserService() *user.Service {
	return p.userService
}
func (p *Provider) GetConfig() *Config {
	return p.cfg
}

func (p *Provider) Shutdown() {
	p.repository.Shutdown()
}
