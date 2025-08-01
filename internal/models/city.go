package models

type City struct {
	ID   int    `db:"city_id" json:"city_id"`
	Name string `db:"name" json:"name"`
}
