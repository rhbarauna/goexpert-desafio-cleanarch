package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/rhbarauna/goexpert-desafio-cleanarch/pkg/events"
	"github.com/streadway/amqp"
)

type OrdersListeddHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewOrdersListeddHandler(rabbitMQChannel *amqp.Channel) *OrdersListeddHandler {
	return &OrdersListeddHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *OrdersListeddHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Orders listed: %v", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish(
		"amq.direct", // exchange
		"",           // key name
		false,        // mandatory
		false,        // immediate
		msgRabbitmq,  // message to publish
	)
}
