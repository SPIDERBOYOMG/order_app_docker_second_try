package models

type Firm struct {
	ID   int    `db:"firm_id" json:"firm_id"`
	Name string `db:"firm_name" json:"firm_namefirm_name"`
}
