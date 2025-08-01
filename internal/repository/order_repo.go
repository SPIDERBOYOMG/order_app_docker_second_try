package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/youruser/order-app/internal/models"
)

type OrderRepo struct {
	db *sqlx.DB
}

func NewOrderRepo(db *sqlx.DB) *OrderRepo {
	return &OrderRepo{
		db: db,
	}
}

func (r *OrderRepo) GetAll(ctx context.Context) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.SelectContext(ctx, &orders, `SELECT * FROM "order"`)

	return orders, err
}

func (r *OrderRepo) GetByID(ctx context.Context, id int) (models.Order, error) {
	var order models.Order
	err := r.db.GetContext(ctx, &order, "SELECT * FROM \"order\" WHERE order_id = $1", id)
	if err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func (r *OrderRepo) Create(ctx context.Context, order models.Order) (int, error) {
	var id int
	err := r.db.GetContext(ctx, &id,
		`INSERT INTO "order" (price, quantity, city_id, firm_id, company_id) VALUES ($1, $2, $3, $4, $5) RETURNING order_id`, order.Price, order.Quantity, order.CityID, order.FirmID, order.CompanyID)

	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *OrderRepo) Update(ctx context.Context, order models.Order) error {
	_, err := r.db.ExecContext(ctx, `UPDATE "order" SET price = $1, quantity = $2, city_id = $3, firm_id = $4, company_id = $5 WHERE order_id = $6`, order.Price, order.Quantity, order.CityID, order.FirmID, order.CompanyID, order.ID)

	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepo) Delete(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM "order" WHERE order_id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
