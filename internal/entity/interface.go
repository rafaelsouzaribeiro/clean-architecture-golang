package entity

type RepositoryOrderInterface interface {
	Save(order *Order) error
	List() ([]Order, error)
}
