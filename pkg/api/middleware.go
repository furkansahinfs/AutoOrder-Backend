package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

// ContextKey implements type for context key
type ContextKey string

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// ContextJWTKey is the key for the jwt context value
const ContextJWTKey ContextKey = "jwt"

// logMiddleware handles logging
func (a *API) logMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.WithFields(logrus.Fields{
			"host":       r.Host,
			"address":    r.RemoteAddr,
			"method":     r.Method,
			"requestURI": r.RequestURI,
			"proto":      r.Proto,
			"useragent":  r.UserAgent(),
		}).Info("HTTP request information")

		next.ServeHTTP(w, r)
	})
}

// corsMiddleware handles preflight
func (a *API) corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Authorization")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		next.ServeHTTP(w, r)
	})
}

// Payload contains the payload data of the token
type Payload struct {
	Email     string    `json:"email"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// NewPayload creates a new token payload with a specific username and duration
func NewPayload(email string, duration time.Duration) (*Payload, error) {
	payload := &Payload{
		Email:     email,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

// Valid checks if the token payload is valid or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

// VerifyToken checks if the token is valid or not
func (a *API) jwtmiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get authentication header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Check if authentication token is present
		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		payload, err := a.verifyToken(authHeaderParts[1], a.config.SigningSecret)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		err = payload.Valid()
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func (a *API) getToken(w http.ResponseWriter, r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("Auth header not found")
	}

	// Check if authentication token is present
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnauthorized)
		return "", errors.New("bearer header not found")

	}
	return authHeaderParts[1], nil

}
func (a *API) verifyToken(token string, signingkey string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(signingkey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		fmt.Println(err)
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
