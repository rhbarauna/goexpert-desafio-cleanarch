package usecase

import (
	"github.com/rhbarauna/goexpert-desafio-cleanarch/internal/entity"
	"github.com/rhbarauna/goexpert-desafio-cleanarch/pkg/events"
)

type OrderListOutputDTO struct {
	Orders []OrderOutputDTO `json:"orders"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrdersListed    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrdersListed events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
		OrdersListed:    OrdersListed,
		EventDispatcher: EventDispatcher,
	}
}

func (c *ListOrdersUseCase) Execute() (OrderListOutputDTO, error) {
	orders, err := c.OrderRepository.List()

	if err != nil {
		return OrderListOutputDTO{}, err
	}

	var outputOrders []OrderOutputDTO

	for _, order := range orders {
		outputOrders = append(
			outputOrders,
			OrderOutputDTO{
				ID:         order.ID,
				Price:      order.Price,
				Tax:        order.Tax,
				FinalPrice: order.FinalPrice,
			},
		)
	}

	dto := OrderListOutputDTO{
		Orders: outputOrders,
	}

	c.OrdersListed.SetPayload(dto)
	c.EventDispatcher.Dispatch(c.OrdersListed)

	return dto, nil
}
