package user_information

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
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
		phone varchar(256) NOT NULL DEFAULT 0,
		address varchar(256) NOT NULL DEFAULT 0,
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
