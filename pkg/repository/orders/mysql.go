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
		orderID varchar(256) NOT NULL DEFAULT '',
		name varchar(256) NOT NULL DEFAULT '',
		date varchar(256) NOT NULL DEFAULT 0,
		brand varchar(256) NOT NULL DEFAULT 0,
		quantity bigint(20) NOT NULL DEFAULT 0,
		price float NOT NULL DEFAULT 0.0,
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
	rows, err := r.db.Query(`SELECT orderID,imagePath, name, brand, quantity, price FROM orders WHERE userID = ?`, userID)

	if err != nil {
		return nil, fmt.Errorf("error get orders: %v", err)
	}

	defer rows.Close()

	orderHistories := make(map[string][]model.Order)

	for rows.Next() {
		var orderHistory model.OrderHistory
		var order model.Order

		err := rows.Scan(&orderHistory.ID, &order.ImagePath, &order.Name, &order.Brand, &order.Quantity, &order.Price)

		if err != nil {
			return nil, fmt.Errorf("error get orders: %v", err)
		}

		orderHistories[orderHistory.ID] = append(orderHistories[orderHistory.ID], order)

	}

	var orderHistoriesSlice []model.OrderHistory

	for index, orderHistory := range orderHistories {
		orderHistoriesSlice = append(orderHistoriesSlice, model.OrderHistory{
			Orders:    orderHistory,
			ImagePath: orderHistory[0].ImagePath,
			ID:        index})

	}

	return orderHistoriesSlice, nil
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
		fmt.Println(order.ImagePath)
		_, err := r.db.Exec(`INSERT INTO orders (userID, orderID, date, brand, quantity, price,imagePath,name) VALUES (?, ?, ?, ?, ?, ?,?,?)`,
			userID, orderID, time.Now().Format("2006-02-01"), order.Brand, order.Quantity, order.Price, order.ImagePath, order.Name)

		if err != nil {
			return fmt.Errorf("error save orders: %v", err)
		}
	}

	return nil
}
