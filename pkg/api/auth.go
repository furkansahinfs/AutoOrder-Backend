package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/furkansahinfs/AutoOrder-Backend/pkg/api/response"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/model"
)

//Login endpoints responsible from user login
func (a *API) Login(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting login info: %v", err), http.StatusBadRequest, a.errors[0].Message)
		return
	}
	dbUser, err := a.service.GetUserService().Login(user, a.config.SigningSecret)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting login info: %v", err), http.StatusBadRequest, a.errors[1].Message)
		return
	}
	response.Write(w, r, dbUser)
	return

}

//Register endpoints responsible from user register
func (a *API) Register(w http.ResponseWriter, r *http.Request) {

}
