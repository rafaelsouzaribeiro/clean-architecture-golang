package database

import (
	"database/sql"
	"testing"

	"github.com/rafaelsouzaribeiro/clean-architecture-golang/internal/entity"
	"github.com/stretchr/testify/suite"

	_ "github.com/mattn/go-sqlite3"
)

type OrderRepositorySuite struct {
	db *sql.DB
	suite.Suite
}

func (o *OrderRepositorySuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	o.NoError(err)
	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	o.db = db
}

func (o *OrderRepositorySuite) TearDownTest() {
	o.db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositorySuite))
}

func (o *OrderRepositorySuite) TestOrderSaveList() {
	order, err := entity.NewOrder("123", 10.0, 2.0)
	o.NoError(err)
	o.NoError(order.CalculateFinalPrice())
	repo := NewOrderRepository(o.db)
	err = repo.Save(order)
	o.NoError(err)

	var orderResult entity.Order
	err = o.db.QueryRow("Select id, price, tax, final_price from orders where id = ?", order.ID).
		Scan(&orderResult.ID, &orderResult.Price, &orderResult.Tax, &orderResult.FinalPrice)

	o.NoError(err)
	o.Equal(order.ID, orderResult.ID)
	o.Equal(order.Price, orderResult.Price)
	o.Equal(order.Tax, orderResult.Tax)
	o.Equal(order.FinalPrice, orderResult.FinalPrice)
}
