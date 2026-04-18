package model

import "time"

type OperationalCost struct {
	ID          int       `db:"id"           json:"id"`
	LocationID  int       `db:"location_id"  json:"location_id"`
	CategoryID  int       `db:"category_id"  json:"category_id"`
	PeriodMonth int8      `db:"period_month" json:"period_month"`
	PeriodYear  int       `db:"period_year"  json:"period_year"`
	Amount      float64   `db:"amount"       json:"amount"`
	Note        *string   `db:"note"         json:"note"`
	CreatedBy   int       `db:"created_by"   json:"created_by"`
	CreatedAt   time.Time `db:"created_at"   json:"created_at"`
}
