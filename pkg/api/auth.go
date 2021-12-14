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
	u, err := a.service.GetUserService().Login(user, a.config.SigningSecret)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting login info: %v", err), http.StatusBadRequest, a.errors[1].Message)
		return
	}
	response.Write(w, r, u)
	return

}

//Login endpoints responsible from user login
func (a *API) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting login info: %v", err), http.StatusBadRequest, a.errors[0].Message)
		return
	}
	_, err = a.verifyToken(user.Token, a.config.SigningSecret)
	if err != nil {
		if err.Error() == "token has expired" {
			u, err := a.service.GetUserService().RefreshToken(&user, a.config.SigningSecret)
			if err != nil {
				response.Errorf(w, r, fmt.Errorf("error getting refresh info: %v", err), http.StatusBadRequest, a.errors[1].Message)
				return
			}
			response.Write(w, r, u)
			return

		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

	}
	response.Write(w, r, user)
	return
}

//Register endpoints responsible from user register
func (a *API) SignUp(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting signup info: %v", err), http.StatusBadRequest, a.errors[2].Message)
		return
	}
	if user.Email == "" {
		response.Errorf(w, r, fmt.Errorf("error getting signup info: %v", err), http.StatusBadRequest, a.errors[2].Message)
		return
	}
	found, err := a.service.GetUserService().CheckExistByMail(user)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting signup info: %v", err), http.StatusBadRequest, a.errors[3].Message)
		return
	}
	if !found {
		u, err := a.service.GetUserService().SignUp(user)
		if err != nil {
			response.Errorf(w, r, fmt.Errorf("error getting signup info: %v", err), http.StatusBadRequest, a.errors[4].Message)
			return
		}
		response.Write(w, r, u)
		return
	}
	response.Errorf(w, r, fmt.Errorf("error getting signup info: %v", err), http.StatusBadRequest, a.errors[5].Message)
	return
}
