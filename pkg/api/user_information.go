package api

import (
	"fmt"
	"net/http"

	"github.com/furkansahinfs/AutoOrder-Backend/pkg/api/response"
)

//endpoint for store userInformation
func (a *API) StoreUserInformation(w http.ResponseWriter, r *http.Request) {
	token, err := a.getToken(w, r)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting storeUserInformation info: %v", err), http.StatusBadRequest, err.Error())
		return
	}

	payload, err := a.verifyToken(token, a.config.SigningSecret)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting storeUserInformation info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	a.service.GetUserService().CheckExistByMail(payload.Email)

	a.service.GetUserInformationService()
}

//endpoint for update userInformation
func (a *API) UpdateUserInformation(w http.ResponseWriter, r *http.Request) {
	token, err := a.getToken(w, r)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting updateUserInformation info: %v", err), http.StatusBadRequest, err.Error())
		return
	}

	payload, err := a.verifyToken(token, a.config.SigningSecret)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting updateUserInformation info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	a.service.GetUserService().CheckExistByMail(payload.Email)
	a.service.GetUserInformationService()
}

//endpoint for delete userInformation
func (a *API) DeleteUserInformation(w http.ResponseWriter, r *http.Request) {
	token, err := a.getToken(w, r)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting deleteUserInformation info: %v", err), http.StatusBadRequest, err.Error())
		return
	}

	payload, err := a.verifyToken(token, a.config.SigningSecret)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting deleteUserInformation info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	a.service.GetUserService().CheckExistByMail(payload.Email)
	a.service.GetUserInformationService()
}
