package router

import (
	"database/sql"

	"github.com/backend-rems/internal/middleware"
	locationHandler "github.com/backend-rems/internal/module/location/handler"
	locationRepository "github.com/backend-rems/internal/module/location/repository"
	locationService "github.com/backend-rems/internal/module/location/service"
	usersHandler "github.com/backend-rems/internal/module/users/handler"
	usersRepository "github.com/backend-rems/internal/module/users/repository"
	usersService "github.com/backend-rems/internal/module/users/service"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	// middelware
	r.Use(middleware.RecoveryMiddleware())
	r.Use(middleware.CORSMiddleware())

	// dependensi injection
	// location
	locationRepo := locationRepository.NewLocationRepository(db)
	locationService := locationService.NewLocationService(locationRepo)
	locationHandler := locationHandler.NewLocationHandler(locationService)

	// user
	usersRepo := usersRepository.NewUsersRepository(db)
	usersService := usersService.NewUsersService(usersRepo)
	usersHandler := usersHandler.NewUsersHandler(usersService)

	api := r.Group("/api/v1")

	api.GET("/provinces", locationHandler.GetAllProvinces)
	api.GET("/provinces/:province_id/regencies", locationHandler.FindRegenciesByProvinceId)
	api.GET("/regencies/:regency_id/districts", locationHandler.GetDistrictByRegencyId)
	api.GET("/districts/:district_id/villages", locationHandler.GetVillageByDistrictId)

	api.POST("/role",usersHandler.InsertRole)

	return r
}
