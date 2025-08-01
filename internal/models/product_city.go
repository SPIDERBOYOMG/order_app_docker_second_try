package models

type ProductCity struct {
	ID int `db:"product_city_id" json:"product_city_id"`

	ProductID int `db:"product_id" json:"product_id"`
	CityID    int `db:"city_id" json:"city_id"`
}
