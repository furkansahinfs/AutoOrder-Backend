package api

import (
	"net/http"

	"github.com/furkansahinfs/AutoOrder-Backend/pkg/api/response"
)

//endpoint for get user configuration
func (a *API) GetUserConfiguration(w http.ResponseWriter, r *http.Request) {

	response.Write(w, r, "information")
}

//endpoint for store user configuration
func (a *API) StoreUserConfiguration(w http.ResponseWriter, r *http.Request) {

	response.Write(w, r, "information")
}

//endpoint for update user configuration
func (a *API) UpdateUserConfiguration(w http.ResponseWriter, r *http.Request) {

	response.Write(w, r, "information")
}

//endpoint for get user configuration
func (a *API) DeleteUserConfiguration(w http.ResponseWriter, r *http.Request) {

	response.Write(w, r, "information")
}
