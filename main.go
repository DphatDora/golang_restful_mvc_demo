package main

import (
	"go_restful_mvc/config"
	"go_restful_mvc/routes"
	"go_restful_mvc/services"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var kafkaProducer *services.KafkaProducer

func main() {
	// Connect to database
	config.ConnectDB()
	config.Migrate()

	// Get Kafka brokers from environment
	kafkaBrokers := os.Getenv("KAFKA_BROKERS")
	if kafkaBrokers == "" {
		kafkaBrokers = "localhost:9092" // Default value
	}
	brokerList := strings.Split(kafkaBrokers, ",")
	log.Printf("Connecting to Kafka brokers: %v", brokerList)

	// Initialize Kafka producer
	var err error
	kafkaProducer, err = services.NewKafkaProducer(brokerList)
	if err != nil {
		log.Printf("Warning: Failed to create Kafka producer: %v", err)
		log.Println("Email notifications will not be sent")
	} else {
		log.Println("Successfully connected to Kafka")
	}

	// Initialize email service
	emailService, err := services.NewEmailService()
	if err != nil {
		log.Printf("Warning: Failed to create email service: %v", err)
		log.Println("Email service will not be available")
	} else {
		log.Println("Email service initialized successfully")
	}

	// Initialize Kafka consumer only if both producer and email service are available
	if kafkaProducer != nil && emailService != nil {
		kafkaConsumer, err := services.NewKafkaConsumer(brokerList, emailService)
		if err != nil {
			log.Printf("Warning: Failed to create Kafka consumer: %v", err)
		} else {
			log.Println("Starting Kafka consumer...")
			kafkaConsumer.ConsumeMessages("user_registered")
		}
	}

	// Set up router
	r := gin.Default()
	routes.RegisterUserRoutes(r, kafkaProducer)
	routes.RegisterProductRoutes(r)

	// Confit static file serving
	r.Static("/user-images", "./user-images")

	// Start the server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on port %s", port)
	r.Run(":" + port)
}
