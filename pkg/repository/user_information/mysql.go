package user_information

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
	tableName = "userInformations"
)

const (
	initTableTemplate = `
	CREATE TABLE IF NOT EXISTS %s (
		id bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
		phone varchar(256) NOT NULL DEFAULT '',
		address varchar(256) NOT NULL DEFAULT '',
		full_name varchar(256) NOT NULL DEFAULT '',
		UNIQUE KEY id (id)
	  ) ENGINE=MyISAM  DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;	
`
)

func NewMySQLRepository(db *sql.DB) (*MySQLRepository, error) {
	tableInitCmd := fmt.Sprintf(initTableTemplate, tableName)
	_, err := db.Exec(tableInitCmd)

	if err != nil {
		return nil, fmt.Errorf("error init userInformations repository: %v", err)
	}

	return &MySQLRepository{
		db: db,
	}, nil
}

func (r *MySQLRepository) GetUserInformation(id int64) (*model.UserInformation, error) {
	q := "SELECT id, phone, address, full_name FROM " + tableName + " WHERE id=?"

	logrus.Debug("QUERY: ", q, "id: ", id)
	res := r.db.QueryRow(q, id)

	information := &model.UserInformation{}

	if err := res.Scan(&information.Id, &information.Phone, &information.Address, &information.FullName); err != nil {
		return nil, err
	}

	return information, nil
}

func (r *MySQLRepository) StoreUserInformation(userInformation model.UserInformation) (int64, error) {
	stmt, err := r.db.Prepare(`INSERT INTO ` + tableName + `(
		phone, address, full_name)
		VALUES(
			?,?,?)`)
	if err != nil {
		return -1, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		userInformation.Phone, userInformation.Address, userInformation.FullName)
	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (r *MySQLRepository) UpdateUserInformation(userInformation model.UserInformation, id int64) (int64, error) {
	result, err := r.db.Exec("update "+tableName+" set phone = ? ,address = ? ,full_name = ?  where id = ?", userInformation.Phone, userInformation.Address, userInformation.FullName, id)
	if err != nil {
		return -1, err
	} else {
		return result.RowsAffected()
	}
}

func (r *MySQLRepository) DeleteUserInformation(id int64) (int64, error) {
	result, err := r.db.Exec("delete from "+tableName+" where id = ?", id)
	if err != nil {
		return -1, err
	} else {
		return result.RowsAffected()
	}
}
