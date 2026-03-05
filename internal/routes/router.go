package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hadinuryani/backend-rems/internal/middleware"
	"go.uber.org/zap"
)

func SetupRouter(logger *zap.Logger) *gin.Engine{
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.RequestLogger(logger))

	api := r.Group("/api")

	registerOfficeRoutes(api)

	return r 
}