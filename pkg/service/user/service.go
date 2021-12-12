package user

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/model"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/repository"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repository repository.Repository
}

func NewService(repo repository.Repository) (*Service, error) {
	return &Service{
		repository: repo,
	}, nil
}

func (s *Service) Login(user model.User, signingKey string) (*model.User, error) {
	//Get user from
	u, err := s.repository.GetUserRepository().GetUser(user)
	if err != nil {
		return nil, err
	}
	if checkPasswordHash(user.Password, u.Password) {
		return nil, errors.New("error when hassing requested User's password")
	}
	u.Password = ""
	claims := jwt.MapClaims{}
	claims["email"] = u.Email
	claims["fullName"] = u.FullName

	token, err := createJWT(claims, signingKey)
	if err != nil {
		return nil, err
	}
	u.Token = token
	return u, nil
}

func (s *Service) SignUp(user model.User) (*model.User, error) {
	u, err := s.repository.GetUserRepository().StoreUser(user)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *Service) CheckExistByMail(user model.User) (bool, error) {
	found, err := s.repository.GetUserRepository().CheckExist(user)
	if err != nil {
		return false, err
	}

	return found, nil
}

//Hash a password and return a string
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//Checks password
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func createJWT(claims jwt.MapClaims, signingSecret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(signingSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
