package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/youruser/order-app/internal/models"
)

type CityRepo struct {
	DB *sqlx.DB
}

func NewCityRepo(db *sqlx.DB) *CityRepo {
	return &CityRepo{
		DB: db,
	}
}

func (r *CityRepo) GetAll(ctx context.Context) ([]models.City, error) {
	var cities []models.City
	err := r.DB.SelectContext(ctx, &cities, "SELECT * FROM city")
	return cities, err
}

func (r *CityRepo) GetByID(ctx context.Context, id int) (models.City, error) {
	var city models.City
	err := r.DB.GetContext(ctx, &city, "SELECT * FROM city WHERE city_id = $1", id)
	if err != nil {
		return models.City{}, err
	}
	return city, nil
}

func (r *CityRepo) Create(ctx context.Context, city models.City) (int, error) {
	var id int
	err := r.DB.QueryRowContext(ctx, "INSERT INTO city (name) VALUES ($1) RETURNING city_id", city.Name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *CityRepo) Update(ctx context.Context, city models.City) error {
	_, err := r.DB.ExecContext(ctx, "UPDATE city SET name = $1 WHERE city_id = $2", city.Name, city.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *CityRepo) Delete(ctx context.Context, id int) error {
	_, err := r.DB.ExecContext(ctx, "DELETE FROM city WHERE city_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
