package models

type Order struct {
	ID       int     `db:"order_id" json:"order_id"`
	Price    float64 `db:"price" json:"price"`
	Quantity int     `db:"quantity" json:"quantity"`

	CityID    int `db:"city_id" json:"city_id"`
	FirmID    int `db:"firm_id" json:"firm_id"`
	CompanyID int `db:"company_id" json:"company_id"`
}
