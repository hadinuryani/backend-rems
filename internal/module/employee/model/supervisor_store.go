package model

import "time"

type SupervisorStore struct {
	SupervisorID int       `db:"supervisor_id" json:"supervisor_id"`
	LocationID   int       `db:"location_id"   json:"location_id"`
	AssignedAt   time.Time `db:"assigned_at"   json:"assigned_at"`
}
