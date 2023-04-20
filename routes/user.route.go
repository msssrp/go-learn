package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/msssrp/go-learn/controllers"
)

func UserRoutes(router *gin.RouterGroup, uc *controllers.UserController) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("", uc.GetAllUsers)
		userGroup.GET("/:id", uc.GetAllUsers)
		userGroup.POST("", uc.CreateUser)
		userGroup.DELETE("/:id", uc.DeleteUser)
	}
}
