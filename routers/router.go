package routers

import (
	"midterm/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoute(route *gin.Engine) {

	apiV1 := route.Group("/api/v1")
	{
		apiV1.POST("role/create", controllers.CreateRole)
	}
}
