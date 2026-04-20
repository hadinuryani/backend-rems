package dto

type ProvinceResponse struct {
	ID           int64  `json:"id"`
	ProvinceName string `json:"province_name"`
}

type RegencyResponse struct {
	ID          int64  `json:"id"`
	RegencyName string `json:"regency_name"`
	ProvinceID  int64  `json:"province_id"`
}

type DistrictResponse struct {
	ID              int64  `json:"id"`
	SubdistrictName string `json:"subdistrict_name"`
	RegencyID       int64  `json:"regency_id"`
}

type VillageResponse struct {
	ID          int64  `json:"id"`
	VillageName string `json:"village_name"`
	DistrictID  int64  `json:"district_id"`
}
