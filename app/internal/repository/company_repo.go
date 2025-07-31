package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/youruser/order-app/internal/models"
)

type CompanyRepo struct {
	db *sqlx.DB
}

func NewCompanyRepo(db *sqlx.DB) *CompanyRepo {
	return &CompanyRepo{
		db: db,
	}
}

func (r *CompanyRepo) GetAll(ctx context.Context) ([]models.Company, error) {
	var companies []models.Company
	err := r.db.SelectContext(ctx, &companies, "SELECT * FROM company")
	return companies, err
}

func (r *CompanyRepo) GetByID(ctx context.Context, id int) (models.Company, error) {
	var company models.Company
	err := r.db.GetContext(ctx, &company, "SELECT * FROM company WHERE company_id = $1", id)
	if err != nil {
		return models.Company{}, err
	}
	return company, nil
}

func (r *CompanyRepo) Create(ctx context.Context, company models.Company) (int, error) {
	var id int
	err := r.db.QueryRowContext(ctx, "INSERT INTO company (name) VALUES ($1) RETURNING company_id", company.Name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *CompanyRepo) Update(ctx context.Context, company models.Company) error {
	_, err := r.db.ExecContext(ctx, "UPDATE company SET company_name = $1 WHERE company_id = $2", company.Name, company.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *CompanyRepo) Delete(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM company WHERE company_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
