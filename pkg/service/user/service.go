package user

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/model"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type Service struct {
	repository repository.Repository
}

func NewService(repo repository.Repository) (*Service, error) {
	return &Service{
		repository: repo,
	}, nil
}

func (s *Service) GetUser(user model.User) (*model.User, error) {
	//Get user from db
	u, err := s.repository.GetUserRepository().GetUser(user)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *Service) Login(user model.User, signingKey string) (*model.User, error) {
	//Get user from
	u, err := s.repository.GetUserRepository().GetUser(user)
	if err != nil {
		return nil, err
	}
	samePassword, err := checkPasswordHash(user.Password, u.Password)
	if err != nil {
		return nil, err
	}
	if !samePassword {

		return nil, errors.New("username or password wrong")
	}
	u.Password = ""
	token, err := CreateToken(u.Email, time.Minute*5000, signingKey)
	if err != nil {
		return nil, err
	}
	u.Token = token
	return u, nil
}

func (s *Service) SignUp(user model.User) (*model.User, error) {
	hash, err := hashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hash
	u, err := s.repository.GetUserRepository().StoreUser(user)
	if err != nil {
		return nil, err
	}
	u.Password = ""
	return u, nil
}

func (s *Service) RefreshToken(user *model.User, signingKey string) (*model.User, error) {
	found, err := s.CheckExistByMail(user.Email)
	if err != nil {
		return nil, err
	}
	if found {
		token, err := CreateToken(user.Email, time.Minute*50000, signingKey)
		if err != nil {
			return nil, err
		}
		user.Token = token
		return user, nil
	}
	return nil, errors.New("User Not Found")

}

func (s *Service) CheckExistByMail(mail string) (bool, error) {
	found, err := s.repository.GetUserRepository().CheckExist(mail)
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
func checkPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, nil
	}
	return true, nil
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

// CreateToken creates a new token for a specific username and duration
func CreateToken(email string, duration time.Duration, signingkey string) (string, error) {
	payload, err := NewPayload(email, duration)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(signingkey))
}

func (s *Service) ChangeUserInformationID(user model.User, id int64) (int64, error) {
	id, err := s.repository.GetUserRepository().ChangeUserInformationID(user, id)
	if err != nil {
		return -1, err
	}
	return id, nil
}
