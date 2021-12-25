package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/furkansahinfs/AutoOrder-Backend/pkg/api/response"
)

//endpoint for get user configuration
func (a *API) GetUserConfiguration(w http.ResponseWriter, r *http.Request) {
	user, err := a.controlUser(w, r)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting storeUserInformation info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	if user.UserInformationID == 0 {
		response.Errorf(w, r, fmt.Errorf("error getting storeUserInformation info: %v", err), http.StatusBadRequest, errors.New("User dont have a information yet").Error())
		return
	}
	information, err := a.service.GetUserInformationService().GetUserInformation(user.UserInformationID)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting storeUserInformation info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	response.Write(w, r, information)
}

//endpoint for store user configuration
func (a *API) StoreUserConfiguration(w http.ResponseWriter, r *http.Request) {
	user, err := a.controlUser(w, r)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting storeUserInformation info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	if user.UserInformationID == 0 {
		response.Errorf(w, r, fmt.Errorf("error getting storeUserInformation info: %v", err), http.StatusBadRequest, errors.New("User dont have a information yet").Error())
		return
	}
	information, err := a.service.GetUserInformationService().GetUserInformation(user.UserInformationID)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting storeUserInformation info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	response.Write(w, r, information)
}

//endpoint for update user configuration
func (a *API) UpdateUserConfiguration(w http.ResponseWriter, r *http.Request) {
	user, err := a.controlUser(w, r)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting storeUserInformation info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	if user.UserInformationID == 0 {
		response.Errorf(w, r, fmt.Errorf("error getting storeUserInformation info: %v", err), http.StatusBadRequest, errors.New("User dont have a information yet").Error())
		return
	}
	information, err := a.service.GetUserInformationService().GetUserInformation(user.UserInformationID)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting storeUserInformation info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	response.Write(w, r, information)
}

//endpoint for get user configuration
func (a *API) DeleteUserConfiguration(w http.ResponseWriter, r *http.Request) {
	user, err := a.controlUser(w, r)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting storeUserInformation info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	if user.UserInformationID == 0 {
		response.Errorf(w, r, fmt.Errorf("error getting storeUserInformation info: %v", err), http.StatusBadRequest, errors.New("User dont have a information yet").Error())
		return
	}
	information, err := a.service.GetUserInformationService().GetUserInformation(user.UserInformationID)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting storeUserInformation info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	response.Write(w, r, information)
}
