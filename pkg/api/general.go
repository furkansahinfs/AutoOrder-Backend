package api

import (
	"errors"
	"net/http"

	"github.com/furkansahinfs/AutoOrder-Backend/pkg/model"
)

func (a *API) controlUser(w http.ResponseWriter, r *http.Request) (*model.User, error) {
	token, err := a.getToken(w, r)
	if err != nil {

		return nil, err
	}
	payload, err := a.verifyToken(token, a.config.SigningSecret)
	if err != nil {
		return nil, err
	}
	found, err := a.service.GetUserService().CheckExistByMail(payload.Email)
	if found {
		var user model.User
		user.Email = payload.Email
		a.service.GetUserService().GetUser(user)
		return &user, nil
	}
	return nil, errors.New("User Not Found")

}