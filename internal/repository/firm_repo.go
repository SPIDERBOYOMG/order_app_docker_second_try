package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/youruser/order-app/internal/models"
)

type FirmRepo struct {
	db *sqlx.DB
}

func NewFirmRepo(db *sqlx.DB) *FirmRepo {
	return &FirmRepo{
		db: db,
	}
}

func (r *FirmRepo) GetAll(ctx context.Context) ([]models.Firm, error) {
	var firms []models.Firm
	err := r.db.SelectContext(ctx, &firms, "SELECT * FROM firm")
	return firms, err
}

func (r *FirmRepo) GetByID(ctx context.Context, id int) (models.Firm, error) {
	var firm models.Firm
	err := r.db.GetContext(ctx, &firm, "SELECT * FROM firm WHERE firm_id = $1", id)
	if err != nil {
		return models.Firm{}, err
	}
	return firm, nil
}

func (r *FirmRepo) Create(ctx context.Context, firm models.Firm) (int, error) {
	var id int
	err := r.db.GetContext(ctx, &id, "INSERT INTO firm (firm_name) VALUES ($1) RETURNING firm_id", firm.Name)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *FirmRepo) Update(ctx context.Context, firm models.Firm) error {
	_, err := r.db.ExecContext(ctx, "UPDATE firm SET firm_name = $1 WHERE firm_id = $2", firm.Name, firm.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *FirmRepo) Delete(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM firm WHERE firm_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
