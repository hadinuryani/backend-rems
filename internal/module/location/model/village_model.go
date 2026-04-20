package model

type Village struct {
	ID          int64  `db:"id"           json:"id"`
	VillageName string `db:"village_name" json:"village_name"`
	DistrictID  int64  `db:"district_id"  json:"district_id"`
}
