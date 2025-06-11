package services

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

type KafkaConsumer struct {
	consumer sarama.Consumer
	emailSvc EmailService
}

type UserRegisteredMessage struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func NewKafkaConsumer(brokers []string, emailSvc EmailService) (*KafkaConsumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		return nil, err
	}
	return &KafkaConsumer{
		consumer: consumer,
		emailSvc: emailSvc,
	}, nil
}

func (c *KafkaConsumer) ConsumeMessages(topic string) {
	partitionConsumer, err := c.consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Printf("Failed to start consumer: %s", err)
		return
	}

	// Create context that can be cancelled
	ctx, cancel := context.WithCancel(context.Background())

	// Handle graceful shutdown
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigterm
		log.Println("Received shutdown signal, closing consumer...")
		cancel()
		// partitionConsumer.Close()
		// c.consumer.Close()
	}()

	go func() {
		defer func() {
			log.Println("Closing partition consumer and main consumer")
			partitionConsumer.Close()
			c.consumer.Close()
		}()

		for {
			select {
			case msg := <-partitionConsumer.Messages():
				var userMsg UserRegisteredMessage
				if err := json.Unmarshal(msg.Value, &userMsg); err != nil {
					log.Printf("Error unmarshaling message: %s", err)
					continue
				}

				// Retry mechanism for sending email
				maxRetries := 3
				for i := 0; i < maxRetries; i++ {
					if err := c.emailSvc.SendWelcomeEmail(userMsg.Email, userMsg.Name); err != nil {
						log.Printf("Error sending welcome email (attempt %d/%d): %s", i+1, maxRetries, err)
						if i < maxRetries-1 {
							continue
						}
					}
					break
				}
			case err := <-partitionConsumer.Errors():
				log.Printf("Error from consumer: %s", err)
			case <-ctx.Done():
				return
			case <-ctx.Done():
				log.Println("Context cancelled, stopping message loop...")
				return
			}

		}
	}()
}
