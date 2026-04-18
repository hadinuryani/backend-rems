package model

import "time"

type Transaction struct {
	ID         int       `db:"id"          json:"id"`
	CategoryID int       `db:"category_id" json:"category_id"`
	LocationID *int      `db:"location_id" json:"location_id"`
	Type       string    `db:"type"        json:"type"` // sale, purchase, salary, operational, distribution, other
	Amount     float64   `db:"amount"      json:"amount"`
	Reference  *string   `db:"reference"   json:"reference"`
	Note       *string   `db:"note"        json:"note"`
	CreatedBy  int       `db:"created_by"  json:"created_by"`
	CreatedAt  time.Time `db:"created_at"  json:"created_at"`
}

