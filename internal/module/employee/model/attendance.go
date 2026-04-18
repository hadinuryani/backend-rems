package model

import "time"

type Attendance struct {
	ID            int        `db:"id"              json:"id"`
	EmployeeID    int        `db:"employee_id"     json:"employee_id"`
	ScheduleID    int        `db:"schedule_id"     json:"schedule_id"`
	CheckIn       *time.Time `db:"check_in"        json:"check_in"`
	CheckOut      *time.Time `db:"check_out"       json:"check_out"`
	CheckInPhoto  *string    `db:"check_in_photo"  json:"check_in_photo"`
	CheckOutPhoto *string    `db:"check_out_photo" json:"check_out_photo"`
	Status        *string    `db:"status"          json:"status"` // present, late, leave, sick, alpha
	Note          *string    `db:"note"            json:"note"`
	CreatedAt     time.Time  `db:"created_at"      json:"created_at"`
}
