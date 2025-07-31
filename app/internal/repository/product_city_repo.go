package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/youruser/order-app/internal/models"
)

type ProductCityRepo struct {
	db *sqlx.DB
}

func NewProductCityRepo(db *sqlx.DB) *ProductCityRepo {
	return &ProductCityRepo{
		db: db,
	}
}

func (r *ProductCityRepo) GetAll(ctx context.Context) ([]models.ProductCity, error) {
	var productCities []models.ProductCity
	err := r.db.SelectContext(ctx, &productCities, "SELECT * FROM product_city")
	return productCities, err
}

func (r *ProductCityRepo) GetByID(ctx context.Context, id int) (models.ProductCity, error) {
	var productCity models.ProductCity
	err := r.db.GetContext(ctx, &productCity, "SELECT * FROM product_city WHERE product_city_id = $1", id)
	if err != nil {
		return models.ProductCity{}, err
	}
	return productCity, nil
}

func (r *ProductCityRepo) Create(ctx context.Context, productCity models.ProductCity) (int, error) {
	var id int
	err := r.db.GetContext(ctx, &id, "INSERT INTO product_city (product_id, city_id) VALUES ($1, $2) RETURNING product_city_id", productCity.ProductID, productCity.CityID)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ProductCityRepo) Update(ctx context.Context, productCity models.ProductCity) error {
	_, err := r.db.ExecContext(ctx, "UPDATE product_city SET product_id = $1, city_id = $2 WHERE product_city_id = $3", productCity.ProductID, productCity.CityID, productCity.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductCityRepo) Delete(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM product_city WHERE product_city_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
