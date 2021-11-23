package user

import (
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

/* func (s *Service) Login(user model.User, signingKey string) (*model.User, error) {
	u, err := s.repository.GetUserRepository().GetUser(user)
	if err != nil {
		return nil, err
	}
}
*/

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
