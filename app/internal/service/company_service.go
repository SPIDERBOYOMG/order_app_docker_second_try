package service

import (
	"context"
	"errors"

	"github.com/youruser/order-app/internal/models"
	"github.com/youruser/order-app/internal/repository"
)

type CompanyService struct {
	Repo *repository.CompanyRepo
}

func NewCompanyService(repo *repository.CompanyRepo) *CompanyService {
	return &CompanyService{
		Repo: repo,
	}
}

func (s *CompanyService) List(ctx context.Context) ([]models.Company, error) {
	companies, err := s.Repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return companies, nil
}

func (s *CompanyService) Get(ctx context.Context, id int) (models.Company, error) {
	if id <= 0 {
		return models.Company{}, errors.New("invalid company ID")
	}
	company, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return models.Company{}, err
	}
	return company, nil
}

func (s *CompanyService) Create(ctx context.Context, company models.Company) (int, error) {
	if company.Name == "" {
		return 0, errors.New("company name cannot be empty")
	}
	id, err := s.Repo.Create(ctx, company)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *CompanyService) Update(ctx context.Context, company models.Company) error {
	if company.ID <= 0 {
		return errors.New("invalid company ID")
	}
	if company.Name == "" {
		return errors.New("company name cannot be empty")
	}
	err := s.Repo.Update(ctx, company)
	if err != nil {
		return err
	}
	return nil
}

func (s *CompanyService) Delete(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("invalid company ID")
	}
	err := s.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
