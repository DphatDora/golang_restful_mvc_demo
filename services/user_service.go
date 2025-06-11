package services

import (
	"go_restful_mvc/models"
	"go_restful_mvc/repositories"
	"log"
	"time"
)

type UserService interface {
	Register(user *models.User) error
	Login(email, password string) (*models.User, error)
	Update(id string, user *models.User) error
}

type userService struct {
	repo  repositories.UserRepository
	kafka *KafkaProducer
}

func NewUserService(repo repositories.UserRepository, kafkaProducer *KafkaProducer) UserService {
	return &userService{
		repo:  repo,
		kafka: kafkaProducer,
	}
}

func (s *userService) Register(user *models.User) error {
	// Lưu user vào database
	if err := s.repo.Create(user); err != nil {
		return err
	}

	// Gửi message đến Kafka trong goroutine riêng
	go func() {
		msg := UserRegisteredMessage{
			Email: user.Email,
			Name:  user.Name,
		}

		// Retry mechanism for Kafka message
		maxRetries := 3
		for i := 0; i < maxRetries; i++ {
			if err := s.kafka.SendMessage("user_registered", msg); err != nil {
				log.Printf("Failed to send message to Kafka (attempt %d/%d): %v", i+1, maxRetries, err)
				if i < maxRetries-1 {
					time.Sleep(time.Second * time.Duration(i+1)) // Exponential backoff
					continue
				}
				// Log error but don't affect the main flow
				log.Printf("Failed to send registration message to Kafka after %d attempts: %v", maxRetries, err)
			}
			break
		}
	}()

	return nil
}

func (s *userService) Login(email, password string) (*models.User, error) {
	return s.repo.FindByEmailAndPassword(email, password)
}

func (s *userService) Update(id string, user *models.User) error {
	return s.repo.Update(id, user)
}
