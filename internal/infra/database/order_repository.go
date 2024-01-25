package database

import (
	"database/sql"

	"github.com/rafaelsouzaribeiro/clean-architecture-golang/internal/entity"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		DB: db,
	}
}

func (o *OrderRepository) Save(order *entity.Order) error {
	stmt, err := o.DB.Prepare("INSERT INTO `order` (id,price,tax,final_price)VALUES(?,?,?,?)")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)

	if err != nil {
		return err
	}

	return nil
}

func (o *OrderRepository) List() ([]entity.Order, error) {
	rows, err := o.DB.Query("SELECT id,price,tax,final_price FROM `order`")

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	orders := []entity.Order{}

	for rows.Next() {
		var id string
		var price, tax, final_price float64

		if err := rows.Scan(&id, &price, &tax, &final_price); err != nil {
			return nil, err
		}

		orders = append(orders, entity.Order{
			ID:         id,
			Price:      price,
			Tax:        tax,
			FinalPrice: final_price,
		})
	}

	return orders, nil
}
