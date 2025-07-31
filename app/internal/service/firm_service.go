package service

import (
	"context"
	"errors"

	"github.com/youruser/order-app/internal/models"
	"github.com/youruser/order-app/internal/repository"
)

type FirmService struct {
	Repo *repository.FirmRepo
}

func NewFirmService(repo *repository.FirmRepo) *FirmService {
	return &FirmService{
		Repo: repo,
	}
}

func (s *FirmService) List(ctx context.Context) ([]models.Firm, error) {
	firms, err := s.Repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return firms, nil
}

func (s *FirmService) Get(ctx context.Context, id int) (models.Firm, error) {
	if id <= 0 {
		return models.Firm{}, errors.New("invalid firm ID")
	}
	firm, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return models.Firm{}, err
	}
	return firm, nil
}

func (s *FirmService) Create(ctx context.Context, firm models.Firm) (int, error) {
	if firm.Name == "" {
		return 0, errors.New("firm name cannot be empty")
	}
	id, err := s.Repo.Create(ctx, firm)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *FirmService) Update(ctx context.Context, firm models.Firm) error {
	if firm.ID <= 0 {
		return errors.New("invalid firm ID")
	}
	if firm.Name == "" {
		return errors.New("firm name cannot be empty")
	}
	err := s.Repo.Update(ctx, firm)
	if err != nil {
		return err
	}
	return nil
}

func (s *FirmService) Delete(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("invalid firm ID")
	}
	err := s.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
