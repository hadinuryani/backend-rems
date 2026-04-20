package service

import (
	"github.com/backend-rems/internal/module/location/model"
	"github.com/backend-rems/internal/module/location/repository"
)

type LocationService interface {
	GetAllProvinces()([]model.Province,error)
	FindRegenciesByProvinceId(provinceId int)([]model.Regency,error)
	GetDistrictByRegencyId(regencyId int)([]model.District,error)
	GetVillageByDistrictId(districtId int)([]model.Village,error)
}

type locationService struct {
	repo repository.LocationRepository
}

func NewLocationService(repo repository.LocationRepository) LocationService{
	return &locationService{repo}
}

func(s *locationService)GetAllProvinces()([]model.Province,error){
	return s.repo.GetAllProvinces()
}

func(s *locationService)FindRegenciesByProvinceId(provinceId int)([]model.Regency,error){
	return s.repo.FindRegenciesByProvinceId(provinceId)
}

func(s *locationService)GetDistrictByRegencyId(regencyId int)([]model.District,error){
	return s.repo.GetDistrictByRegencyId(regencyId)
}

func(s *locationService)GetVillageByDistrictId(districtId int)([]model.Village,error){
	return s.repo.GetVillageByDistrictId(districtId)
}