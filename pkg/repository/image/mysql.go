package image

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
	tableName = "images"
)

const (
	initTableTemplate = `
	CREATE TABLE IF NOT EXISTS %s (
		id bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
		user_id bigint(20) NOT NULL DEFAULT 0,
		path varchar(256) NOT NULL DEFAULT '',
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

func (r *MySQLRepository) GetImagePath(id int64) (*model.ImageData, error) {
	q := "SELECT id, path, user_id FROM " + tableName + " WHERE user_id=?"

	logrus.Debug("QUERY: ", q, "user_id: ", id)
	res := r.db.QueryRow(q, id)

	image := &model.ImageData{}

	if err := res.Scan(&image.Id, &image.Path, &image.UserID, id); err != nil {
		return nil, err
	}

	return image, nil
}

func (r *MySQLRepository) StoreImagePath(path string, user_id int64) (int64, error) {
	stmt, err := r.db.Prepare(`INSERT INTO ` + tableName + `(
		path, user_id)
		VALUES(
			?,?)`)
	if err != nil {
		return -1, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		path, user_id)
	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}
