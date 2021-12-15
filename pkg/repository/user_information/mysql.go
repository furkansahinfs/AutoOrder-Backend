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
