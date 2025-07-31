package service

import (
	"context"
	"errors"

	"github.com/youruser/order-app/internal/models"
	"github.com/youruser/order-app/internal/repository"
)

type OrderService struct {
	Repo *repository.OrderRepo
}

func NewOrderService(repo *repository.OrderRepo) *OrderService {
	return &OrderService{
		Repo: repo,
	}
}

func (s *OrderService) List(ctx context.Context) ([]models.Order, error) {
	orders, err := s.Repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *OrderService) Get(ctx context.Context, id int) (models.Order, error) {
	if id <= 0 {
		return models.Order{}, errors.New("invalid order ID")
	}
	order, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func (s *OrderService) Create(ctx context.Context, order models.Order) (int, error) {
	if order.Price <= 0 || order.Quantity <= 0 {
		return 0, errors.New("order price and quantity must be greater than zero")
	}
	id, err := s.Repo.Create(ctx, order)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *OrderService) Update(ctx context.Context, order models.Order) error {
	if order.ID <= 0 {
		return errors.New("invalid order ID")
	}
	if order.Price <= 0 || order.Quantity <= 0 {
		return errors.New("order price and quantity must be greater than zero")
	}
	err := s.Repo.Update(ctx, order)
	if err != nil {
		return err
	}
	return nil
}

func (s *OrderService) Delete(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("invalid order ID")
	}
	err := s.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
