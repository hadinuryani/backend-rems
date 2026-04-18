package model

import "time"

type CashRegister struct {
	ID              int        `db:"id"               json:"id"`
	LocationID      int        `db:"location_id"      json:"location_id"`
	ShiftID         int        `db:"shift_id"         json:"shift_id"`
	EmployeeID      int        `db:"employee_id"      json:"employee_id"`
	OpeningBalance  float64    `db:"opening_balance"  json:"opening_balance"`
	ClosingBalance  *float64   `db:"closing_balance"  json:"closing_balance"`
	ExpectedBalance *float64   `db:"expected_balance" json:"expected_balance"`
	Difference      *float64   `db:"difference"       json:"difference"`
	OpenedAt        time.Time  `db:"opened_at"        json:"opened_at"`
	ClosedAt        *time.Time `db:"closed_at"        json:"closed_at"`
	Status          string     `db:"status"           json:"status"` // open, closed
}
