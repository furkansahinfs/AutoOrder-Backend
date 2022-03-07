package repository

import (
	"database/sql"

	"github.com/furkansahinfs/AutoOrder-Backend/pkg/repository/configuration"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/repository/image"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/repository/orders"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/repository/user"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/repository/user_information"
	_ "github.com/go-sql-driver/mysql"
)

// MySQL Repository defines the MySQL implementation of Repository interface
type MySQLRepository struct {
	cfg                 *MySQLConfig
	db                  *sql.DB
	userRepo            user.Repository
	userInformationRepo user_information.Repository
	imageRepo           image.Repository
	configurationRepo   configuration.Repository
	OrdersRepo          orders.Repository
}

// MySQLConfig defines the MySQL Repository configuration
type MySQLConfig struct {
	Addr     string `yaml:"addr"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
}

// dbConn opens connection with MySQL driver
func dbConn(cfg *MySQLConfig) (*sql.DB, error) {

	dbDriver := "mysql"    // Database driver
	dbUser := cfg.Username // Mysql username
	dbPass := cfg.Password // Mysql password
	dbName := cfg.DBName   // Mysql schema
	addr := cfg.Addr

	// Realize the connection with mysql driver
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+addr+")/")
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		return nil, err
	}

	db.Close()

	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=true")
	if err != nil {
		return nil, err
	}

	// Return db object to be used by other functions
	return db, nil
}

// NewMySQLRepository creates a new MySQL Repository
func NewMySQLRepository(cfg *MySQLConfig) (*MySQLRepository, error) {
	db, err := dbConn(cfg)
	if err != nil {
		return nil, err
	}
	userRepo, err := user.NewMySQLRepository(db)
	if err != nil {
		return nil, err
	}
	userInformationRepo, err := user_information.NewMySQLRepository(db)
	if err != nil {
		return nil, err
	}
	imageRepo, err := image.NewMySQLRepository(db)
	if err != nil {
		return nil, err
	}
	configurationRepo, err := configuration.NewMySQLRepository(db)
	if err != nil {
		return nil, err
	}
	ordersRepo, err := orders.NewMySQLRepository(db)
	if err != nil {
		return nil, err
	}
	return &MySQLRepository{
		cfg:                 cfg,
		db:                  db,
		userRepo:            userRepo,
		userInformationRepo: userInformationRepo,
		imageRepo:           imageRepo,
		configurationRepo:   configurationRepo,
		OrdersRepo:          ordersRepo,
	}, nil
}

// GetUserRepository returns the user repository
func (r *MySQLRepository) GetUserRepository() user.Repository {
	return r.userRepo
}

// GetUserInformationRepository returns the user repository
func (r *MySQLRepository) GetUserInformationRepository() user_information.Repository {
	return r.userInformationRepo
}

// GetImageRepository returns the user repository
func (r *MySQLRepository) GetImageRepository() image.Repository {
	return r.imageRepo
}

// GetConfigurationRepository returns the configuration repository
func (r *MySQLRepository) GetConfigurationRepository() configuration.Repository {
	return r.configurationRepo
}

// GetOrdersRepository returns the orders repository
func (r *MySQLRepository) GetOrdersRepository() orders.Repository {
	return r.OrdersRepo
}

// Shutdown closes the database connection
func (r *MySQLRepository) Shutdown() {
	r.db.Close()
}
