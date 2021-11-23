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
	a.service.GetUserService()
}

//Register endpoints responsible from user register
func (a *API) Register(w http.ResponseWriter, r *http.Request) {

}
