package server

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/furkansahinfs/AutoOrder-Backend/pkg/api"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/model"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/repository"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/service"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Config struct {
	ListenAddress string `yaml:"listen_address"`
	CertFile      string `yaml:"cert_file"`
	KeyFile       string `yaml:"key_file"`

	API           *api.Config             `yaml:"api"`
	Service       *service.Config         `yaml:"service"`
	DB            *repository.MySQLConfig `yaml:"database"`
	ErrorFilePath string                  `yaml:"errorFilePath"`
}

// Instance represents an instance of the server
type Instance struct {
	Config     *Config
	API        *api.API
	Service    service.Service
	httpServer *http.Server
}

// NewInstance returns an new instance of our server
func NewInstance(cfg *Config) *Instance {
	return &Instance{
		Config: cfg,
	}
}

// Start starts the server
func (i *Instance) Start() {
	var err error
	var router = mux.NewRouter()

	// Establish database connection

	repo, err := repository.NewMySQLRepository(i.Config.DB)
	if err != nil {
		logrus.WithError(err).Fatal("Could not create mysql repository")
	}

	i.Service, err = service.NewProvider(i.Config.Service, repo)
	if err != nil {
		logrus.WithError(err).Fatal("Could not create service provider")
	}

	var errJson model.ErrorJson
	file, err := ioutil.ReadFile(i.Config.ErrorFilePath)
	if err != nil {
		logrus.WithError(err).Fatal("Error not found errors_en.json")
		return
	}
	err = json.Unmarshal(file, &errJson)
	if err != nil {
		logrus.WithError(err).Fatal("Error parsing errors_en.json")
		return
	}

	// Initialize API
	i.API, err = api.New(i.Config.API, i.Service, router, errJson)
	if err != nil {
		logrus.WithError(err).Fatal("Could not create API instance")
	}
	// Startup the HTTP Server in a way that we can gracefully shut it down again
	i.httpServer = &http.Server{
		Addr:    i.Config.ListenAddress,
		Handler: router,
	}

	err = i.httpServer.ListenAndServe()
	if err != http.ErrServerClosed {
		logrus.WithError(err).Error("HTTP Server stopped unexpected")
		i.Shutdown()
	} else {
		logrus.WithError(err).Info("HTTP Server stopped")
	}
}

// Shutdown stops the server
func (i *Instance) Shutdown() {
	// Shutdown all dependencies
	i.Service.Shutdown()

	// Shutdown HTTP server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := i.httpServer.Shutdown(ctx)
	if err != nil {
		logrus.WithError(err).Error("Failed to shutdown HTTP server gracefully")
		os.Exit(1)
	}

	logrus.Info("Shutdown HTTP server...")
	os.Exit(0)
}
