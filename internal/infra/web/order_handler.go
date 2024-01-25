package web

import (
	"encoding/json"
	"net/http"

	"github.com/rafaelsouzaribeiro/clean-architecture-golang/internal/entity"
	"github.com/rafaelsouzaribeiro/clean-architecture-golang/internal/usecase"
	"github.com/rafaelsouzaribeiro/clean-architecture-golang/pkg/events"
)

type WebOrderHandle struct {
	EventDispatcher   events.EventDispatcherInterface
	OrderRepository   entity.RepositoryOrderInterface
	OrderCreatedEvent events.EventInterface
}

func NewWebOrderHandle(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.RepositoryOrderInterface,
	OrderCreatedEvent events.EventInterface,
) *WebOrderHandle {
	return &WebOrderHandle{
		EventDispatcher:   EventDispatcher,
		OrderRepository:   OrderRepository,
		OrderCreatedEvent: OrderCreatedEvent,
	}
}

func (web *WebOrderHandle) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.InputOrderDto

	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createOrder := usecase.NewCreateOrderUseCase(web.OrderRepository, web.OrderCreatedEvent, web.EventDispatcher)
	output, err := createOrder.Execute(dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *WebOrderHandle) List(w http.ResponseWriter, r *http.Request) {
	listOrder := usecase.NewListOrderUseCase(h.OrderRepository)

	output, err := listOrder.Execute()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
