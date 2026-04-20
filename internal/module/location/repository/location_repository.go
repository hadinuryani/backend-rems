package repository

import (
	"database/sql"

	"github.com/backend-rems/internal/module/location/model"
)

type LocationRepository interface {
	GetAllProvinces()([]model.Province,error)
	FindRegenciesByProvinceId(provinceId int)([]model.Regency,error)
	GetDistrictByRegencyId(regencyId int)([]model.District,error)
	GetVillageByDistrictId(districtId int)([]model.Village,error)
}

type locationRepository struct {
	db *sql.DB
}

func NewLocationRepository(db *sql.DB)LocationRepository{
	return &locationRepository{db}
}

func (r *locationRepository)GetAllProvinces()([]model.Province,error){
	rows,err := r.db.Query("SELECT id, province_name FROM provinces")
	if err != nil {
		return nil,err
	}
	defer rows.Close()

	var provices []model.Province
	for rows.Next() {
		var prov model.Province
		err := rows.Scan(&prov.ID,&prov.ProvinceName)
		if err != nil {
			return nil,err
		}
		provices = append(provices, prov)
	}
	return provices,nil
}

func (r *locationRepository)FindRegenciesByProvinceId(provinceId int)([]model.Regency,error){
	rows,err := r.db.Query("SELECT id,regency_name,province_id FROM regency WHERE province_id = ?",provinceId)
	if err != nil{
		return nil,err
	}
	defer rows.Close()
	
	var regencys []model.Regency
	for rows.Next(){
		var reg model.Regency
		err := rows.Scan(&reg.ID,&reg.RegencyName,&reg.ProvinceID)
		if err != nil{
			return nil,err
		}
		regencys = append(regencys, reg)
	}
	return regencys,nil
}

func(r *locationRepository) GetDistrictByRegencyId(regencyId int)([]model.District,error){
	rows,err := r.db.Query("SELECT id,subdistrict_name,regency_id FROM districts WHERE regency_id = ?",regencyId)
	if err != nil {
		return nil,err
	}
	defer rows.Close()

	var districts []model.District
	for rows.Next(){
		var district model.District
		err := rows.Scan(&district.ID,&district.SubdistrictName,&district.RegencyID)
		if err != nil {
			return nil,err
		}
		districts = append(districts, district)
	}
	return districts,nil
}

func(r *locationRepository)GetVillageByDistrictId(districtId int)([]model.Village,error){
	rows,err := r.db.Query("SELECT id,village_name,district_id FROM village WHERE district_id = ?",districtId)
	if err != nil{
		return nil,err
	}
	defer rows.Close()

	var villages []model.Village
	for rows.Next(){
		var village model.Village
		err := rows.Scan(&village.ID,&village.VillageName,&village.DistrictID)
		if err != nil{
			return nil,err
		}
		villages = append(villages, village)
	}
	return villages,nil
}