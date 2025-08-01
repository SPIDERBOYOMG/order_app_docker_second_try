package models

type Company struct {
	ID   int    `db:"company_id" json:"company_id"`
	Name string `db:"company_name" json:"company_name"`
}
