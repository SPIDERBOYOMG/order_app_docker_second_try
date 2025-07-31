package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/youruser/order-app/internal/models"
)

type ProductRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) *ProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (r *ProductRepo) GetAll(ctx context.Context) ([]models.Product, error) {
	var products []models.Product
	err := r.db.SelectContext(ctx, &products, "SELECT * FROM product")
	return products, err
}

func (r *ProductRepo) GetByID(ctx context.Context, id int) (models.Product, error) {
	var product models.Product
	err := r.db.GetContext(ctx, &product, "SELECT * FROM product WHERE product_id = $1", id)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func (r *ProductRepo) Create(ctx context.Context, product models.Product) (int, error) {
	var id int
	err := r.db.GetContext(ctx, &id, "INSERT INTO product (product_name, firm_id) VALUES ($1, $2) RETURNING product_id", product.Name, product.FirmID)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ProductRepo) Update(ctx context.Context, product models.Product) error {
	_, err := r.db.ExecContext(ctx, "UPDATE product SET product_name = $1, firm_id = $2 WHERE product_id = $3", product.Name, product.FirmID, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepo) Delete(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM product WHERE product_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
