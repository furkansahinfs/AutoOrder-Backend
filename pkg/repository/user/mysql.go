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
		user_information_id bigint(20) NOT NULL DEFAULT 0,
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

func (r *MySQLRepository) CheckExist(user model.User) (bool, error) {
	var exists bool
	row := r.db.QueryRow(`SELECT EXISTS(SELECT 1 FROM `+tableName+` WHERE email=? )`, user.Email)
	if err := row.Scan(&exists); err != nil {
		return false, err
	}
	return exists, nil
}

func (r *MySQLRepository) StoreUser(user model.User) (*model.User, error) {
	stmt, err := r.db.Prepare(`INSERT INTO ` + tableName + `(
		email, password)
		VALUES(
			?,?,?)`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		user.Email, user.Password, user.FullName)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
