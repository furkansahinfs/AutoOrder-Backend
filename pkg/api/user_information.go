package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/furkansahinfs/AutoOrder-Backend/pkg/api/response"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/model"
)

//endpoint for store userInformation
func (a *API) StoreUserInformation(w http.ResponseWriter, r *http.Request) {
	user, err := a.controlUser(w, r)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting updateUserInformation info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	var payload model.UserInformation
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting login info: %v", err), http.StatusBadRequest, err.Error())
		return
	}

	a.service.GetUserInformationService()
}

//endpoint for update userInformation
func (a *API) UpdateUserInformation(w http.ResponseWriter, r *http.Request) {
	user, err := a.controlUser(w, r)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting updateUserInformation info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	var payload model.UserInformation
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting login info: %v", err), http.StatusBadRequest, err.Error())
		return
	}

	a.service.GetUserInformationService()
}

//endpoint for delete userInformation
func (a *API) DeleteUserInformation(w http.ResponseWriter, r *http.Request) {
	user, err := a.controlUser(w, r)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting updateUserInformation info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	var payload model.UserInformation
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting login info: %v", err), http.StatusBadRequest, err.Error())
		return
	}

	a.service.GetUserInformationService()
}
