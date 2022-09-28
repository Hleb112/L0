package service

import (
	"github.com/google/uuid"
	"log"
	cache "testgo/internal/cache"
	"testgo/internal/models"
	"testgo/internal/repository"
	"time"
)

type Service struct {
	repo  *repository.Repository
	cache *cache.Cache
}

func New(repo *repository.Repository, cache *cache.Cache) *Service {
	return &Service{
		repo:  repo,
		cache: cache,
	}
}

func (s Service) SaveOrder(order *models.Order) error {
	err := s.repo.SaveOrder(order)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) GetOrderByUUID(uuid uuid.UUID) (*models.Order, error) {
	orderCache, exist := s.cache.Get(uuid)
	if exist {
		log.Println("got from cache")
		return orderCache, nil
	}

	order, err := s.repo.GetOrderByUuid(uuid)
	if err != nil {
		return nil, err
	}

	s.cache.Set(uuid, order, 5*time.Minute)

	return order, nil
}

func (s Service) GetOrders() ([]models.Order, error) {
	orders, err := s.repo.GetOrders()
	if err != nil {
		return []models.Order{}, err
	}

	return orders, nil
}
