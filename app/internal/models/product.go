package models

type Product struct {
	ID   int    `db:"product_id" json:"product_id"`
	Name string `db:"product_name" json:"product_name"`

	FirmID int `db:"firm_id" json:"firm_id"`
}
