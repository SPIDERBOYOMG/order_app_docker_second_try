package service

import (
	"context"
	"errors"

	"github.com/youruser/order-app/internal/models"
	"github.com/youruser/order-app/internal/repository"
)

type ProductCityService struct {
	Repo *repository.ProductCityRepo
}

func NewProductCityService(repo *repository.ProductCityRepo) *ProductCityService {
	return &ProductCityService{
		Repo: repo,
	}
}

func (s *ProductCityService) List(ctx context.Context) ([]models.ProductCity, error) {
	productCities, err := s.Repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return productCities, nil
}

func (s *ProductCityService) Get(ctx context.Context, id int) (models.ProductCity, error) {
	if id <= 0 {
		return models.ProductCity{}, errors.New("invalid product city ID")
	}
	productCity, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return models.ProductCity{}, err
	}
	return productCity, nil
}

func (s *ProductCityService) Create(ctx context.Context, productCity models.ProductCity) (int, error) {
	id, err := s.Repo.Create(ctx, productCity)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *ProductCityService) Update(ctx context.Context, productCity models.ProductCity) error {
	if productCity.ID <= 0 {
		return errors.New("invalid product city ID")
	}
	err := s.Repo.Update(ctx, productCity)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductCityService) Delete(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("invalid product city ID")
	}
	err := s.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
