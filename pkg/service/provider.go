package service

import "github.com/furkansahinfs/AutoOrder-Backend/pkg/repository"

type Provider struct {
	cfg        *Config
	repository repository.Repository
}

func NewProvider(cfg *Config, repo repository.Repository) (*Provider, error) {

	return &Provider{
		cfg:        cfg,
		repository: repo,
	}, nil
}

func (p *Provider) GetConfig() *Config {
	return p.cfg
}

func (p *Provider) Shutdown() {
	p.repository.Shutdown()
}
