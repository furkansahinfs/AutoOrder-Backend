package user

import (
	"database/sql"
	"fmt"

	"github.com/furkansahinfs/AutoOrder-Backend/pkg/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

type MySQLRepository struct {
	db *sql.DB
}

const (
	tableName = "users"
)

const (
	initTableTemplate = `
	CREATE TABLE IF NOT EXISTS %s (
		id bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
		email varchar(256) NOT NULL DEFAULT 0,
		password varchar(256) NOT NULL DEFAULT 0,
		full_name varchar(256) NOT NULL DEFAULT 0,
		token varchar(256) NOT NULL DEFAULT '',
		UNIQUE KEY id (id)
	  ) ENGINE=MyISAM  DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;	
`
)

func NewMySQLRepository(db *sql.DB) (*MySQLRepository, error) {
	tableInitCmd := fmt.Sprintf(initTableTemplate, tableName)
	_, err := db.Exec(tableInitCmd)

	if err != nil {
		return nil, fmt.Errorf("error init user repository: %v", err)
	}

	return &MySQLRepository{
		db: db,
	}, nil
}

func (r *MySQLRepository) GetUser(user model.User) (*model.User, error) {
	q := "SELECT email, password FROM " + tableName + " WHERE email=?"

	logrus.Debug("QUERY: ", q, "email: ", user.Email)
	res := r.db.QueryRow(q, user.Email)

	u := &model.User{}

	if err := res.Scan(&u.Email, &u.Password); err != nil {
		return nil, err
	}

	return u, nil
}
