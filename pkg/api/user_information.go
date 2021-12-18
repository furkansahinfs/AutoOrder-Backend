package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/furkansahinfs/AutoOrder-Backend/pkg/api/response"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/model"
)

//endpoint for store userInformation
func (a *API) StoreUserInformation(w http.ResponseWriter, r *http.Request) {
	user, err := a.controlUser(w, r)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting storeUserInformation info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	var payload model.UserInformation
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting storeUserInformation info: %v", err), http.StatusBadRequest, err.Error())
		return
	}

	if user.UserInformationID != 0 {
		response.Errorf(w, r, fmt.Errorf("error getting storeUserInformation info: %v", err), http.StatusBadRequest, errors.New("User have a information already").Error())
		return
	}
	id, err := a.service.GetUserInformationService().StoreUserInformation(payload)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting storeUserInformation info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	_, err = a.service.GetUserService().ChangeUserInformationID(*user, id)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting storeUserInformation info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	response.Write(w, r, user)
}

//endpoint for update userInformation
func (a *API) UpdateUserInformation(w http.ResponseWriter, r *http.Request) {
	user, err := a.controlUser(w, r)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting UpdateUserInformation info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	var payload model.UserInformation
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting UpdateUserInformation info: %v", err), http.StatusBadRequest, err.Error())
		return
	}

	if user.UserInformationID == 0 {
		response.Errorf(w, r, fmt.Errorf("error getting UpdateUserInformation info: %v", err), http.StatusBadRequest, errors.New("User dont have any information yet").Error())
		return
	}
	_, err = a.service.GetUserInformationService().UpdateUserInformation(payload, user.UserInformationID)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting UpdateUserInformation info: %v", err), http.StatusBadRequest, err.Error())
		return
	}

	response.Write(w, r, user)
}

//endpoint for delete userInformation
func (a *API) DeleteUserInformation(w http.ResponseWriter, r *http.Request) {
	user, err := a.controlUser(w, r)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting DeleteUserInformation info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	var payload model.UserInformation
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting DeleteUserInformation info: %v", err), http.StatusBadRequest, err.Error())
		return
	}

	if user.UserInformationID != 0 {
		response.Errorf(w, r, fmt.Errorf("error getting DeleteUserInformation info: %v", err), http.StatusBadRequest, errors.New("User dont have any information yet").Error())
		return
	}
	_, err = a.service.GetUserInformationService().DeleteUserInformation(user.UserInformationID)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting DeleteUserInformation info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	_, err = a.service.GetUserService().ChangeUserInformationID(*user, 0)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting DeleteUserInformation info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	response.Write(w, r, user)
}
