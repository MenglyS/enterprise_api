package routers

import (
	"midterm/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoute(route *gin.Engine) {

	apiV1 := route.Group("/api/v1")
	{
		apiV1.POST("role/create", controllers.CreateRole)

		//language routes
		apiV1.POST("language/create", controllers.CreateLanguage)
		apiV1.GET("language/getAll", controllers.GetLanguages)
		apiV1.PUT("language/edit/:id", controllers.EditLanguage)
		apiV1.DELETE("language/delete/:id", controllers.DeleteLanguage)

		//position routes
		apiV1.POST("position/create", controllers.CreatePosition)
		apiV1.GET("position/getAll", controllers.GetPositions)
		apiV1.PUT("position/edit/:id", controllers.EditPosition)
		apiV1.DELETE("position/delete/:id", controllers.DeletePosition)

		//skill routes
		apiV1.POST("skill/create", controllers.CreateSkill)
		apiV1.GET("skill/getAll", controllers.GetSkills)
		apiV1.PUT("skill/edit/:id", controllers.EditSkill)
		apiV1.DELETE("skill/delete/:id", controllers.DeleteSkill)

		//category routes
		apiV1.POST("category/create", controllers.CreateCategory)
		apiV1.GET("category/getAll", controllers.GetCategorys)
		apiV1.PUT("category/edit/:id", controllers.EditCategory)
		apiV1.DELETE("category/delete/:id", controllers.DeleteCategory)

		//user routes
		apiV1.POST("user/create", controllers.CreateUser)
		apiV1.GET("user/getAll", controllers.GetUsers)
		apiV1.PUT("user/edit/:id", controllers.EditUser)
		apiV1.DELETE("user/delete/:id", controllers.DeleteUser)

		//job routes
		apiV1.POST("job/create", controllers.CreateJob)
		apiV1.GET("job/getAll", controllers.GetJobs)
		apiV1.DELETE("job/delete/:id", controllers.DeleteJob)
		apiV1.PUT("job/edit/:id", controllers.EditJob)

		//job routes
		apiV1.POST("applicant/create/:id", controllers.CreateApplicant)
		apiV1.GET("applicant/getAll", controllers.GetApplicants)
		apiV1.DELETE("applicant/delete/:id", controllers.DeleteApplicant)
		apiV1.PUT("applicant/edit/:id", controllers.EditApplicant)
	}
}
