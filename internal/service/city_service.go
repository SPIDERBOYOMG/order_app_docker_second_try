package service

import (
	"context"
	"errors"

	"github.com/youruser/order-app/internal/models"
	"github.com/youruser/order-app/internal/repository"
)

type CityService struct {
	Repo *repository.CityRepo
}

func NewCityService(repo *repository.CityRepo) *CityService {
	return &CityService{
		Repo: repo,
	}
}

func (s *CityService) List(ctx context.Context) ([]models.City, error) {
	cities, err := s.Repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return cities, nil
}

func (s *CityService) Get(ctx context.Context, id int) (models.City, error) {
	if id <= 0 {
		return models.City{}, errors.New("invalid city ID")
	}
	city, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return models.City{}, err
	}
	return city, nil
}

func (s *CityService) Create(ctx context.Context, city models.City) (int, error) {
	if city.Name == "" {
		return 0, errors.New("city name cannot be empty")
	}
	id, err := s.Repo.Create(ctx, city)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *CityService) Update(ctx context.Context, city models.City) error {
	if city.ID <= 0 {
		return errors.New("invalid city ID")
	}
	if city.Name == "" {
		return errors.New("city name cannot be empty")
	}
	err := s.Repo.Update(ctx, city)
	if err != nil {
		return err
	}
	return nil
}

func (s *CityService) Delete(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("invalid city ID")
	}
	err := s.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
