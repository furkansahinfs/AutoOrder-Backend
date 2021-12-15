package repository

import (
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/repository/user"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/repository/user_information"
)

// Repository defines the method for all operations related with repository
// Repository interface is composition of  Repository interfaces of imported packages.
type Repository interface {
	Shutdown()
	GetUserRepository() user.Repository
	GetUserInformationRepository() user_information.Repository
}
