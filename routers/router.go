package routers

import (
	"midterm/controllers"
	"midterm/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute(route *gin.Engine) {

	apiV1 := route.Group("/api/v1")
	{
		apiV1.POST("role/create", middleware.AuthMiddleware(), controllers.CreateRole)

		//admin routes
		admin := apiV1.Group("/admin")
		admin.POST("login", controllers.AdminLogin)

		//language routes
		language := apiV1.Group("/language")
		language.Use(middleware.AuthMiddleware())
		language.POST("/create", controllers.CreateLanguage)
		language.GET("/getAll", controllers.GetLanguages)
		language.PUT("/edit/:id", controllers.EditLanguage)
		language.DELETE("/delete/:id", controllers.DeleteLanguage)

		//position routes
		position := apiV1.Group("/position")
		position.POST("/create", controllers.CreatePosition)
		position.GET("/getAll", controllers.GetPositions)
		position.PUT("/edit/:id", controllers.EditPosition)
		position.DELETE("/delete/:id", controllers.DeletePosition)

		//skill routes
		skill := apiV1.Group("/skill")
		skill.POST("/create", controllers.CreateSkill)
		skill.GET("/getAll", controllers.GetSkills)
		skill.PUT("/edit/:id", controllers.EditSkill)
		skill.DELETE("/delete/:id", controllers.DeleteSkill)

		//category routes
		category := apiV1.Group("/category")
		category.POST("/create", controllers.CreateCategory)
		category.GET("/getAll", controllers.GetCategorys)
		category.PUT("/edit/:id", controllers.EditCategory)
		category.DELETE("/delete/:id", controllers.DeleteCategory)

		//user routes
		user := apiV1.Group("/user")
		user.POST("/create", controllers.CreateUser)
		user.GET("/getAll", controllers.GetUsers)
		user.PUT("/edit/:id", controllers.EditUser)
		user.DELETE("/delete/:id", controllers.DeleteUser)

		//job routes
		job := apiV1.Group("/job")
		job.POST("/create", controllers.CreateJob)
		job.GET("/getAll", controllers.GetJobs)
		job.DELETE("/delete/:id", controllers.DeleteJob)
		job.PUT("/edit/:id", controllers.EditJob)

		//applicants routes
		applicant := apiV1.Group("/applicant")
		applicant.POST("/create/:id", controllers.CreateApplicant)
		applicant.GET("/getAll", controllers.GetApplicants)
		applicant.DELETE("/delete/:id", controllers.DeleteApplicant)
		applicant.PUT("/edit/:id", controllers.EditApplicant)
		applicant.GET("/count", controllers.CountApplicants)

		//expense routes
		expense := apiV1.Group("/expense")
		expense.POST("/create", controllers.CreateExpense)
		expense.GET("/getAll", controllers.GetExpenses)
		expense.PUT("/edit/:id", controllers.EditExpense)
		expense.DELETE("/delete/:id", controllers.DeleteExpense)

		//leave routes
		leave := apiV1.Group("/leave")
		leave.POST("/create", controllers.CreateLeave)
		leave.GET("/getAll", controllers.GetLeaves)
		leave.PUT("/edit/:id", controllers.EditLeave)
		leave.DELETE("/delete/:id", controllers.DeleteLeave)
	}
}
