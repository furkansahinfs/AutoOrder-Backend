package api

import (
	"net/http"

	"github.com/furkansahinfs/AutoOrder-Backend/pkg/model"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/service"
	"github.com/gorilla/mux"
	_ "github.com/mitchellh/mapstructure"
)

// API configuration
type Config struct {
	Domain        string `yaml:"domain"`
	SigningSecret string `yaml:"signing_secret"`
	ImagePath     string `yaml:"imagePath"`
	Enums         struct {
		Front []map[string]string `yaml:"front"`
		Back  []map[string]string `yaml:"back"`
	} `yaml:"enum"`
	PythonBackendAddress string `yaml:"python_backend_address"`
	FakeApiAddress       string `yaml:"fake_api_address"`
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
	api.Router.HandleFunc("/api/v1/userinformation", api.corsMiddleware(api.logMiddleware(api.jwtmiddleware(api.GetUserInformation)))).Methods("GET")

	//image endpoints
	api.Router.HandleFunc("/api/v1/image", api.corsMiddleware(api.logMiddleware(api.jwtmiddleware(api.GetImage)))).Methods("POST")

	//config endpoints
	api.Router.HandleFunc("/api/v1/configuration/{type}", api.corsMiddleware(api.logMiddleware(api.jwtmiddleware(api.GetUserConfiguration)))).Methods("GET")
	api.Router.HandleFunc("/api/v1/configuration/store/{type}", api.corsMiddleware(api.logMiddleware(api.jwtmiddleware(api.StoreUserConfiguration)))).Methods("POST")
	api.Router.HandleFunc("/api/v1/configuration/delete/{type}", api.corsMiddleware(api.logMiddleware(api.jwtmiddleware(api.DeleteUserConfiguration)))).Methods("POST")
	api.Router.HandleFunc("/api/v1/configuration/update/{type}", api.corsMiddleware(api.logMiddleware(api.jwtmiddleware(api.UpdateUserConfiguration)))).Methods("POST")

	//enum endpoints
	api.Router.HandleFunc("/api/v1/items/front", api.corsMiddleware(api.logMiddleware(api.jwtmiddleware(api.GetItemsFront)))).Methods("GET")
	api.Router.HandleFunc("/api/v1/items/back", api.corsMiddleware(api.logMiddleware(api.jwtmiddleware(api.GetItemsBack)))).Methods("GET")
	api.Router.HandleFunc("/api/v1/items/all", api.corsMiddleware(api.logMiddleware(api.jwtmiddleware(api.GetItemsAll)))).Methods("GET")

	//order history endpoints
	api.Router.HandleFunc("/api/v1/orderhistory", api.corsMiddleware(api.logMiddleware(api.jwtmiddleware(api.GetOrderHistory)))).Methods("GET")

	// healtcheck endpoint
	api.Router.HandleFunc("/api/v1/healtcheck", api.corsMiddleware(api.logMiddleware(api.jwtmiddleware(api.preflightHandler)))).Methods("POST")

	//image serve
	s := http.StripPrefix("/api/v1/images/", http.FileServer(http.Dir("./pkg/images/")))
	api.Router.PathPrefix("/api/v1/images/").Handler(s)

	return api, nil

}

func (a *API) preflightHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	return
}
