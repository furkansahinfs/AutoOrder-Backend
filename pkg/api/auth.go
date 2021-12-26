package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/api/response"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/model"
)

//Login endpoints responsible from user login
func (a *API) Login(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting login info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	u, err := a.service.GetUserService().Login(user, a.config.SigningSecret)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting login info: %v", err), http.StatusBadRequest, err.Error())
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
			token, _, err := new(jwt.Parser).ParseUnverified(user.Token, jwt.MapClaims{})
			if err != nil {
				log.Fatal(err)
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				log.Fatalf("Can't convert token's claims to standard claims")
			}
			switch email := claims["email"].(type) {
			case string:
				user.Email = email
			}

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
	found, err := a.service.GetUserService().CheckExistByMail(user.Email)
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

func (a *API) extractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecretString := a.config.SigningSecret
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}
