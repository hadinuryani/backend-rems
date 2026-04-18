package model

import "time"

type Employee struct {
	ID            int        `db:"id"            json:"id"`
	RoleID        *int       `db:"role_id"       json:"role_id"`
	StatusID      *int       `db:"status_id"     json:"status_id"`
	LocationID    *int       `db:"location_id"   json:"location_id"`
	ManagedBy     *int       `db:"managed_by"    json:"managed_by"`
	VillageID     *int64     `db:"village_id"    json:"village_id"`
	Name          string     `db:"name"          json:"name"`
	Email         string     `db:"email"         json:"email"`
	NoHp          string     `db:"no_hp"         json:"no_hp"`
	NIK           string     `db:"nik"           json:"nik"`
	AddressDetail *string    `db:"address_detail" json:"address_detail"`
	Photo         *string    `db:"photo"         json:"photo"`
	UpdatedAt     *time.Time `db:"updated_at"    json:"updated_at"`
	CreatedAt     time.Time  `db:"created_at"    json:"created_at"`
}