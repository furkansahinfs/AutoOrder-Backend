package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/furkansahinfs/AutoOrder-Backend/pkg/api/response"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/model"
	"github.com/gorilla/mux"
)

//endpoint for get user configuration
func (a *API) GetUserConfiguration(w http.ResponseWriter, r *http.Request) {
	user, err := a.controlUser(w, r)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting user configuration info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	vars := mux.Vars(r)
	switch vars["type"] {
	case "front":
		items, err := a.service.GetUserConfigurationService().GetUserConfiguration(user.ID, vars["type"])
		if err != nil {
			response.Errorf(w, r, fmt.Errorf("error getting user configuration info: %v", err), http.StatusBadRequest, err.Error())
			return
		}
		response.Write(w, r, items)
		return
	case "back":
		items, err := a.service.GetUserConfigurationService().GetUserConfiguration(user.ID, vars["type"])
		if err != nil {
			response.Errorf(w, r, fmt.Errorf("error getting user configuration info: %v", err), http.StatusBadRequest, err.Error())
			return
		}
		response.Write(w, r, items)
		return
	default:
		response.Errorf(w, r, fmt.Errorf("error getting user configuration info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
}

//endpoint for store user configuration
func (a *API) StoreUserConfiguration(w http.ResponseWriter, r *http.Request) {
	user, err := a.controlUser(w, r)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting user configuration info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	vars := mux.Vars(r)
	switch vars["type"] {
	case "front":
		var uitems []model.Item
		err := json.NewDecoder(r.Body).Decode(&uitems)
		if err != nil {
			response.Errorf(w, r, fmt.Errorf("error getting user configuration info: %v", err), http.StatusBadRequest, a.errors[2].Message)
			return
		}
		frontItems, err := a.GetItems()
		if err != nil {
			response.Errorf(w, r, fmt.Errorf("error getting user configuration info: %v", err), http.StatusBadRequest, err.Error())
			return
		}
		err = a.service.GetUserConfigurationService().StoreUserConfiguration(user.ID, vars["type"], uitems, frontItems)
		if err != nil {
			response.Errorf(w, r, fmt.Errorf("error getting user configuration info: %v", err), http.StatusBadRequest, err.Error())
			return
		}
		response.Write(w, r, true)
		return
	case "back":
		var uitems []model.Item
		err := json.NewDecoder(r.Body).Decode(&uitems)
		if err != nil {
			response.Errorf(w, r, fmt.Errorf("error getting user configuration info: %v", err), http.StatusBadRequest, a.errors[2].Message)
			return
		}
		backItems, err := a.GetItems()
		if err != nil {
			response.Errorf(w, r, fmt.Errorf("error getting user configuration info: %v", err), http.StatusBadRequest, err.Error())
			return
		}
		err = a.service.GetUserConfigurationService().StoreUserConfiguration(user.ID, vars["type"], uitems, backItems)
		if err != nil {
			response.Errorf(w, r, fmt.Errorf("error getting user configuration info: %v", err), http.StatusBadRequest, err.Error())
			return
		}
		response.Write(w, r, true)
		return
	default:
		response.Errorf(w, r, fmt.Errorf("error getting user configuration info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
}

//endpoint for update user configuration
func (a *API) UpdateUserConfiguration(w http.ResponseWriter, r *http.Request) {
	user, err := a.controlUser(w, r)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting user configuration info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	vars := mux.Vars(r)
	switch vars["type"] {
	case "front":
		var uitems []model.Item
		err := json.NewDecoder(r.Body).Decode(&uitems)
		if err != nil {
			response.Errorf(w, r, fmt.Errorf("error getting user configuration info: %v", err), http.StatusBadRequest, a.errors[2].Message)
			return
		}
		frontItems, err := a.GetItems()
		if err != nil {
			response.Errorf(w, r, fmt.Errorf("error getting user configuration info: %v", err), http.StatusBadRequest, err.Error())
			return
		}
		err = a.service.GetUserConfigurationService().UpdateUserConfiguration(user.ID, vars["type"], uitems, frontItems)
		if err != nil {
			response.Errorf(w, r, fmt.Errorf("error getting user configuration info: %v", err), http.StatusBadRequest, err.Error())
			return
		}
		response.Write(w, r, true)
		return
	case "back":
		var uitems []model.Item
		err := json.NewDecoder(r.Body).Decode(&uitems)
		if err != nil {
			response.Errorf(w, r, fmt.Errorf("error getting user configuration info: %v", err), http.StatusBadRequest, a.errors[2].Message)
			return
		}
		backItems, err := a.GetItems()
		if err != nil {
			response.Errorf(w, r, fmt.Errorf("error getting user configuration info: %v", err), http.StatusBadRequest, err.Error())
			return
		}
		err = a.service.GetUserConfigurationService().UpdateUserConfiguration(user.ID, vars["type"], uitems, backItems)
		if err != nil {
			response.Errorf(w, r, fmt.Errorf("error getting user configuration info: %v", err), http.StatusBadRequest, err.Error())
			return
		}
		response.Write(w, r, true)
		return
	default:
		response.Errorf(w, r, fmt.Errorf("error getting user configuration info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
}

//endpoint for get user configuration
func (a *API) DeleteUserConfiguration(w http.ResponseWriter, r *http.Request) {
	user, err := a.controlUser(w, r)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting user configuration info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	vars := mux.Vars(r)
	switch vars["type"] {
	case "front":

		err = a.service.GetUserConfigurationService().DeleteUserConfiguration(user.ID, vars["type"])
		if err != nil {
			response.Errorf(w, r, fmt.Errorf("error getting user configuration info: %v", err), http.StatusBadRequest, err.Error())
			return
		}
		response.Write(w, r, true)
		return
	case "back":
		err = a.service.GetUserConfigurationService().DeleteUserConfiguration(user.ID, vars["type"])
		if err != nil {
			response.Errorf(w, r, fmt.Errorf("error getting user configuration info: %v", err), http.StatusBadRequest, err.Error())
			return
		}
		response.Write(w, r, true)
		return
	default:
		response.Errorf(w, r, fmt.Errorf("error getting user configuration info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
}
