package orders

import (
	"database/sql"
	"fmt"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/model"
	"time"
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
		userID bigint(20) NOT NULL DEFAULT 0,
		orderID bigint(20) NOT NULL DEFAULT 0,
		name varchar(256) NOT NULL DEFAULT '',
		date varchar(256) NOT NULL DEFAULT 0,
		brand varchar(256) NOT NULL DEFAULT 0,
		quantity bigint(20) NOT NULL DEFAULT 0,
		price bigint(20) NOT NULL DEFAULT 0,
		imagePath varchar(256) NOT NULL DEFAULT 0,
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

func (r *MySQLRepository) GetOrdersWithGroupByOrderID(userID int64) ([]model.OrderHistory, error) {
	rows, err := r.db.Query(`SELECT orderID,imagePath,name, brand, quantity, price FROM orders WHERE userID = ? GROUP BY orderID`, userID)

	if err != nil {
		return nil, fmt.Errorf("error get orders: %v", err)
	}

	defer rows.Close()

	var orderHistories []model.OrderHistory

	for rows.Next() {
		var orderHistory model.OrderHistory
		var order model.Order

		err := rows.Scan(&orderHistory.ID, &orderHistory.ImagePath, &order.Name, &order.Brand, &order.Quantity, &order.Price)

		if err != nil {
			return nil, fmt.Errorf("error get orders: %v", err)
		}

		for _, orderH := range orderHistories {
			if orderH.ID == orderHistory.ID {
				orderHistory.Orders = append(orderHistory.Orders, order)
			} else {
				orderHistories = append(orderHistories, orderHistory)
			}
		}
	}

	return orderHistories, nil
}

func (r *MySQLRepository) GetOrder(userID int64, orderID string) (model.OrderHistory, error) {
	rows, err := r.db.Query(`SELECT orderID,imagePath,name, brand, quantity, price FROM orders WHERE userID = ? AND orderID = ?`, userID, orderID)

	if err != nil {
		return model.OrderHistory{}, fmt.Errorf("error get orders: %v", err)
	}

	defer rows.Close()

	var orderHistory model.OrderHistory

	for rows.Next() {
		var order model.Order

		err := rows.Scan(&orderHistory.ID, &orderHistory.ImagePath, &order.Name, &order.Brand, &order.Quantity, &order.Price)

		if err != nil {
			return model.OrderHistory{}, fmt.Errorf("error get orders: %v", err)
		}

		orderHistory.Orders = append(orderHistory.Orders, order)
	}

	return orderHistory, nil
}

func (r *MySQLRepository) SaveOrder(order []model.OrderResponse, userID int64, orderID string) error {
	for _, order := range order {
		_, err := r.db.Exec(`INSERT INTO orders (userID, orderID, date, brand, quantity, price,imagePath) VALUES (?, ?, ?, ?, ?, ?,?)`,
			userID, orderID, time.Now().Format("2006-02-01"), order.Brand, order.Quantity, order.Price, order.ImagePath)

		if err != nil {
			return fmt.Errorf("error save orders: %v", err)
		}
	}

	return nil
}
