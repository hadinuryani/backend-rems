package routes

import "github.com/gin-gonic/gin"

func registerOfficeRoutes(rg *gin.RouterGroup) {
	office := rg.Group("/office")
	{
		office.GET("")
	}
}