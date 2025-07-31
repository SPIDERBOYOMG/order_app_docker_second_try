package service

import (
	"context"
	"errors"

	"github.com/youruser/order-app/internal/models"
	"github.com/youruser/order-app/internal/repository"
)

type ProductService struct {
	Repo *repository.ProductRepo
}

func NewProductService(repo *repository.ProductRepo) *ProductService {
	return &ProductService{
		Repo: repo,
	}
}

func (s *ProductService) List(ctx context.Context) ([]models.Product, error) {
	products, err := s.Repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) Get(ctx context.Context, id int) (models.Product, error) {
	if id <= 0 {
		return models.Product{}, errors.New("invalid product ID")
	}
	product, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func (s *ProductService) Create(ctx context.Context, product models.Product) (int, error) {
	if product.Name == "" {
		return 0, errors.New("product name cannot be empty")
	}
	id, err := s.Repo.Create(ctx, product)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *ProductService) Update(ctx context.Context, product models.Product) error {
	if product.ID <= 0 {
		return errors.New("invalid product ID")
	}
	if product.Name == "" {
		return errors.New("product name cannot be empty")
	}
	err := s.Repo.Update(ctx, product)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductService) Delete(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("invalid product ID")
	}
	err := s.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
