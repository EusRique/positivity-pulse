package main

import (
	"net/http"
	"positivitypulse/internal/adapters"
	"positivitypulse/web"
	"time"

	"github.com/gin-gonic/gin"
)

func StartScheduler() {
	ticker := time.NewTicker(1 * time.Second) // Ajuste para o tempo desejado
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			adapters.FetchAdviceAndPublish()
		}
	}
}

func main() {
	adapters.InitRabbitMQ()
	adapters.InitPostgres()

	defer adapters.CloseRabbitMQ()

	r := gin.Default()

	r.GET("/alive", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Its live!")
	})

	r.GET("/fetch-advice", func(c *gin.Context) {
		adapters.FetchAdviceAndPublish()
		c.JSON(200, gin.H{"message": "Message sent to the queue"})
	})

	r.GET("/ws", func(c *gin.Context) {
		web.HandleConnections(c.Writer, c.Request)
	})

	go adapters.ConsumeMessages()
	//go StartScheduler()

	r.Run(":3000")
}
