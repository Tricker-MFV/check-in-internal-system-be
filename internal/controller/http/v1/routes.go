package v1

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func MapRoutes(router *gin.Engine, employeeHandler *EmployeeHandler) {
	v1 := router.Group("/api")
	{
		employees := v1.Group("/employees")
		{
			employees.GET("/", employeeHandler.GetAll)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
