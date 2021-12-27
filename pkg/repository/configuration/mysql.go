package configuration

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
	tableName = "configurations"
)

const (
	initTableTemplate = `
	CREATE TABLE IF NOT EXISTS %s (
		id bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
		item_name varchar(256) NOT NULL DEFAULT '',
		item_size varchar(256) NOT NULL DEFAULT '',
		item_type varchar(256) NOT NULL DEFAULT '',
		user_id bigint(20) NOT NULL DEFAULT 0,
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

func (r *MySQLRepository) GetConfiguration(id int64, item_type string) ([]model.Item, error) {
	q := "SELECT item_name, item_size, item_type FROM " + tableName + " WHERE user_id=? AND item_type=?"

	logrus.Debug("QUERY: ", q, "id: item_type", id)
	res, err := r.db.Query(q, id, item_type)
	if err != nil {
		return nil, fmt.Errorf("error init userInformations repository: %v", err)
	}
	var items []model.Item
	for res.Next() {
		var item model.Item
		if err := res.Scan(&item.Name, &item.Size, &item.Type); err != nil {
			return nil, err
		}
		items = append(items, item)

	}

	return items, nil
}

func (r *MySQLRepository) StoreConfiguration(items []model.Item, user_id int64) error {
	for _, item := range items {
		stmt, err := r.db.Prepare(`INSERT INTO ` + tableName + `(
			item_name, item_type, item_size, user_id)
			VALUES(
				?,?,?,?)`)
		if err != nil {
			return err
		}

		defer stmt.Close()

		_, err = stmt.Exec(
			item.Name, item.Type, item.Size, user_id)
		if err != nil {
			return err
		}

	}
	return nil

}

func (r *MySQLRepository) UpdateConfiguration(id int64, items []model.Item, item_type string) error {
	err := r.DeleteConfiguration(id, item_type)
	if err != nil {
		return err
	}
	err = r.StoreConfiguration(items, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *MySQLRepository) DeleteConfiguration(id int64, item_type string) error {
	_, err := r.db.Exec("delete from "+tableName+" where user_id = ? AND item_type = ?", id, item_type)
	if err != nil {
		return err
	} else {
		return nil
	}
}
