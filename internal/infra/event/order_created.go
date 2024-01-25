package event

import "time"

type OrderCreated struct {
	Name    string
	Payload interface{}
}

func NewOrderCreated() *OrderCreated {
	return &OrderCreated{
		Name: "OrderCreated",
	}
}

func (o *OrderCreated) GetName() string {
	return o.Name
}

func (o *OrderCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *OrderCreated) SetPayload(payload interface{}) {
	o.Payload = payload
}

func (e *OrderCreated) GetDateTime() time.Time {
	return time.Now()
}
