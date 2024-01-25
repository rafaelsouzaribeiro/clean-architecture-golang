package usecase

import "github.com/rafaelsouzaribeiro/clean-architecture-golang/internal/entity"

type ListOrderUseCase struct {
	OrderRepository entity.RepositoryOrderInterface
}

func NewListOrderUseCase(OrderRepository entity.RepositoryOrderInterface) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (l *ListOrderUseCase) Execute() ([]OutputOrderDto, error) {
	orders, err := l.OrderRepository.List()

	if err != nil {
		return []OutputOrderDto{}, err
	}

	outPutDto := []OutputOrderDto{}

	for _, order := range orders {
		outPutDto = append(outPutDto, OutputOrderDto{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}
	return outPutDto, nil
}
