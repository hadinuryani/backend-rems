package model

type District struct {
	ID              int64  `db:"id"               json:"id"`
	SubdistrictName string `db:"subdistrict_name" json:"subdistrict_name"`
	RegencyID       int64  `db:"regency_id"       json:"regency_id"`
}