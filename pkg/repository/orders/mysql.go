package orders

import (
	"database/sql"
	"fmt"

	"github.com/furkansahinfs/AutoOrder-Backend/pkg/model"
)

type MySQLRepository struct {
	db *sql.DB
}

const (
	tableName = "orders"
)

const (
	initTableTemplate = `
	CREATE TABLE IF NOT EXISTS %s (
		id bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
		orderGroupID bigint(20) NOT NULL DEFAULT 0,
		userID bigint(20) NOT NULL DEFAULT 0,
		date varchar(256) NOT NULL DEFAULT 0,
		ItemName bigint(20) NOT NULL DEFAULT 0,
		count bigint(20) NOT NULL DEFAULT 0,
		UNIQUE KEY id (id)
	  ) ENGINE=MyISAM  DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;	
`
)

func NewMySQLRepository(db *sql.DB) (*MySQLRepository, error) {
	tableInitCmd := fmt.Sprintf(initTableTemplate, tableName)
	_, err := db.Exec(tableInitCmd)

	if err != nil {
		return nil, fmt.Errorf("error init orders repository: %v", err)
	}

	return &MySQLRepository{
		db: db,
	}, nil
}

func (r *MySQLRepository) GetOrders(userID int64) ([]*model.Order, error) {
	rows, err := r.db.Query("SELECT * FROM orders where userID=?", userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		err := rows.Scan(&order.ID, &order.UserID, &order.Date, &order.ItemName, &order.Count)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (r *MySQLRepository) GetOrder(orderGroupID int64, userID int64) (*model.Order, error) {
	q := "SELECT * FROM orders WHERE orderGroupID=? AND userID=?"
	res := r.db.QueryRow(q, orderGroupID, userID)

	u := &model.Order{}

	if err := res.Scan(&u.ID, &u.UserID, &u.Date, &u.ItemName, &u.Count); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *MySQLRepository) CreateOrder(order *model.Order, userID int64) error {
	q := "INSERT INTO orders (userID, date, itemName, count) VALUES (?, ?, ?, ?)"
	_, err := r.db.Exec(q, userID, order.Date, order.ItemName, order.Count)

	return err
}
func (r *MySQLRepository) GetlastOrder() (*model.Order, error) {
	q := "SELECT * FROM orders ORDER BY id DESC LIMIT 1"
	res := r.db.QueryRow(q)

	u := &model.Order{}

	if err := res.Scan(&u.ID, &u.UserID, &u.Date, &u.ItemName, &u.Count); err != nil {
		return nil, err
	}

	return u, nil
}
