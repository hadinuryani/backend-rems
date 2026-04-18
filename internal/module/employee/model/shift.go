package model

import "time"

type Shift struct {
	ID        int       `db:"id"         json:"id"`
	ShiftName string    `db:"shift_name" json:"shift_name"`
	StartTime string    `db:"start_time" json:"start_time"` // TIME format HH:MM:SS
	EndTime   string    `db:"end_time"   json:"end_time"`   // TIME format HH:MM:SS
	CreatedBy *int      `db:"created_by" json:"created_by"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
