package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/backend-rems/internal/module/location/dto"
	"github.com/backend-rems/internal/module/location/service"
	"github.com/backend-rems/pkg/response"
	"github.com/gin-gonic/gin"
)

type LocationHandler struct {
	service service.LocationService
}

func NewLocationHandler(s service.LocationService) *LocationHandler {
	return &LocationHandler{
		service: s,
	}
}

func (h *LocationHandler) GetAllProvinces(c *gin.Context) {
	province, err := h.service.GetAllProvinces()
	var res []dto.ProvinceResponse
	if err != nil {
		log.Println(err.Error())
		response.ErrorResponse(c, http.StatusInternalServerError, "gagal mengambil data provinsi", nil)
		return
	}

	for _, prov := range province {
		res = append(res, dto.ProvinceResponse{
			ID:           prov.ID,
			ProvinceName: prov.ProvinceName,
		})
	}
	response.SuccesResponse(c, http.StatusOK, "sukses mengambil data provinsi", res)
}

func (h *LocationHandler) FindRegenciesByProvinceId(c *gin.Context) {
	provinceId := c.Param("province_id")
	idInt, err := strconv.Atoi(provinceId)
	if err != nil {
		log.Println(err.Error())
		response.ErrorResponse(c, http.StatusBadRequest, "id provinsi tidak valid", nil)
		return
	}

	regencys, err := h.service.FindRegenciesByProvinceId(idInt)
	var res []dto.RegencyResponse
	for _, regency := range regencys {
		res = append(res, dto.RegencyResponse{
			ID:          regency.ID,
			RegencyName: regency.RegencyName,
			ProvinceID:  regency.ProvinceID,
		})
	}
	response.SuccesResponse(c, http.StatusOK, "sukses mengambil data kabupaten", res)
}

func (h *LocationHandler) GetDistrictByRegencyId(c *gin.Context) {
	regencyId := c.Param("regency_id")
	idInt, err := strconv.Atoi(regencyId)
	if err != nil {
		log.Println(err.Error())
		response.ErrorResponse(c, http.StatusBadRequest, "id kabupaten tidak valid", nil)
		return
	}
	districts, err := h.service.GetDistrictByRegencyId(idInt)
	var res []dto.DistrictResponse
	for _, district := range districts {
		res = append(res, dto.DistrictResponse{
			ID:              district.ID,
			SubdistrictName: district.SubdistrictName,
			RegencyID:       district.RegencyID,
		})
	}
	response.SuccesResponse(c, http.StatusOK, "sukses mengambil data kecamatan", res)
}

func (h *LocationHandler) GetVillageByDistrictId(c *gin.Context) {
	districtId := c.Param("district_id")
	IdInt, err := strconv.Atoi(districtId)
	if err != nil {
		log.Println(err.Error())
		response.ErrorResponse(c, http.StatusBadRequest, "id kecamatan tidak valid", nil)
		return
	}

	villages, err := h.service.GetVillageByDistrictId(IdInt)
	var res []dto.VillageResponse
	for _, village := range villages {
		res = append(res, dto.VillageResponse{
			ID:          village.ID,
			VillageName: village.VillageName,
			DistrictID:  village.DistrictID,
		})
	}
	response.SuccesResponse(c,http.StatusOK,"sukses mengambil data desa",res)
}
