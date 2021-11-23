package repository

import "github.com/furkansahinfs/AutoOrder-Backend/pkg/repository/user"

// Repository defines the method for all operations related with repository
// Repository interface is composition of  Repository interfaces of imported packages.
type Repository interface {
	Shutdown()
	GetUserRepository() user.Repository
}
