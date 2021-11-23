package api

import (
	"net/http"

	"github.com/furkansahinfs/AutoOrder-Backend/pkg/model"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/service"
	"github.com/gorilla/mux"
)

// API configuration
type Config struct {
	Domain        string `yaml:"domain"`
	SigningSecret string `yaml:"signing_secret"`
}

// structure of the API
type API struct {
	Router *mux.Router

	config  *Config
	service service.Service
	errors  model.ErrorJson
}

// New returns the api settings
func New(config *Config, svc service.Service, router *mux.Router, errors model.ErrorJson) (*API, error) {
	api := &API{
		config:  config,
		service: svc,
		Router:  router,
		errors:  errors,
	}

	// Endpoint for browser preflight requests
	api.Router.Methods("OPTIONS").HandlerFunc(api.corsMiddleware(api.preflightHandler))

	// auth endpoints
	api.Router.HandleFunc("/api/v1/login", api.corsMiddleware(api.logMiddleware(api.Login))).Methods("POST")
	//api.Router.HandleFunc("/api/v1/signup", api.corsMiddleware(api.logMiddleware(api.SignUP))).Methods("POST")
	return api, nil

}

func (a *API) preflightHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	return
}
