package model

import "time"

type Role struct {
	ID          int       `db:"id"          json:"id"`
	RoleName    string    `db:"role_name"   json:"role_name"`
	Level       *int8     `db:"level"       json:"level"` // 1=manager, 2=hrd/admin, 3=supervisor, 4=kepala, 5=helper
	Scope       *string   `db:"scope"       json:"scope"` // office, store, warehouse, all
	Description *string   `db:"description" json:"description"`
	IsActive    bool      `db:"is_active"   json:"is_active"`
	CreatedAt   time.Time `db:"created_at"  json:"created_at"`
}
