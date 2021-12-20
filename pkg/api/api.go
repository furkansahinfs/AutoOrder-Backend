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
	ImagePath     string `yaml:"imagePath"`
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
	api.Router.HandleFunc("/api/v1/signup", api.corsMiddleware(api.logMiddleware(api.SignUp))).Methods("POST")
	api.Router.HandleFunc("/api/v1/refreshtoken", api.corsMiddleware(api.logMiddleware(api.RefreshToken))).Methods("POST")

	//user information endpoints
	api.Router.HandleFunc("/api/v1/userinformation/update", api.corsMiddleware(api.logMiddleware(api.jwtmiddleware(api.UpdateUserInformation)))).Methods("POST")
	api.Router.HandleFunc("/api/v1/userinformation/delete", api.corsMiddleware(api.logMiddleware(api.jwtmiddleware(api.DeleteUserInformation)))).Methods("POST")
	api.Router.HandleFunc("/api/v1/userinformation/store", api.corsMiddleware(api.logMiddleware(api.jwtmiddleware(api.StoreUserInformation)))).Methods("POST")

	//image endpoints
	api.Router.HandleFunc("/api/v1/image", api.corsMiddleware(api.logMiddleware(api.jwtmiddleware(api.GetImage)))).Methods("POST")

	// healtcheck endpoint
	api.Router.HandleFunc("/api/v1/healtcheck", api.corsMiddleware(api.logMiddleware(api.jwtmiddleware(api.preflightHandler)))).Methods("POST")
	return api, nil

}

func (a *API) preflightHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	return
}
