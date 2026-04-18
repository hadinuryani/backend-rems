package model

import "time"

type Schedule struct {
	ID            int       `db:"id"             json:"id"`
	EmployeeID    int       `db:"employee_id"    json:"employee_id"`
	ShiftID       int       `db:"shift_id"       json:"shift_id"`
	ScheduledDate time.Time `db:"scheduled_date" json:"scheduled_date"`
	CreatedBy     int       `db:"created_by"     json:"created_by"`
	CreatedAt     time.Time `db:"created_at"     json:"created_at"`
}
