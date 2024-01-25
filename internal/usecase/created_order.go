package usecase

import (
	"github.com/google/uuid"
	"github.com/rafaelsouzaribeiro/clean-architecture-golang/internal/entity"
	"github.com/rafaelsouzaribeiro/clean-architecture-golang/pkg/events"
)

type InputOrderDto struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

type OutputOrderDto struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_Price"`
}

type CreateOrderUseCase struct {
	OrderRepository entity.RepositoryOrderInterface
	OrderCreated    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewCreateOrderUseCase(
	OrderRepository entity.RepositoryOrderInterface,
	OrderCreated events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderRepository: OrderRepository,
		OrderCreated:    OrderCreated,
		EventDispatcher: EventDispatcher,
	}
}

func (c *CreateOrderUseCase) Execute(input InputOrderDto) (OutputOrderDto, error) {
	order, err := entity.NewOrder(uuid.NewString(), input.Price, input.Tax)

	if err != nil {
		return OutputOrderDto{}, err
	}

	err = order.CalculateFinalPrice()

	if err != nil {
		return OutputOrderDto{}, err
	}

	if err := c.OrderRepository.Save(order); err != nil {
		return OutputOrderDto{}, err
	}

	dto := OutputOrderDto{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}

	c.OrderCreated.SetPayload(dto)
	c.EventDispatcher.Dispatch(c.OrderCreated)

	return dto, nil
}
