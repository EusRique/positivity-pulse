package adapters

import (
	"encoding/json"
	"log"

	"positivitypulse/web"
)

type AdviceMessage struct {
	ID     int    `json:"id"`
	Advice string `json:"advice"`
}

func ConsumeMessages() {
	ch := GetRabbitMQChannel()
	if ch == nil {
		log.Println("Error: RabbitMQ channel is not available")
		return
	}

	msgs, err := ch.Consume("advice_queue", "", true, false, false, false, nil)
	if err != nil {
		log.Printf("Error consuming messages: %v", err)
		return
	}

	for msg := range msgs {
		var advice AdviceMessage
		err := json.Unmarshal(msg.Body, &advice)
		if err != nil {
			log.Printf("Error processing message: %v", err)
			continue
		}

		log.Println("Saving data to the database...")
		DB.Create(&Advice{Message: advice.Advice})

		log.Println("Sent message to frontend via WebSocket")
		web.NotifyFrontend(advice.Advice)
	}
}
