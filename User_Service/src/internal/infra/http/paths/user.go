package paths

import (
	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/infra/http/handlers"
	"github.com/gin-gonic/gin"
)

func AddUserRoutes(router *gin.Engine) {
	userRoutes := router.Group("/user")
	userRoutes.POST("/", handlers.ImpNewUserHandler.CreateUser)
	userRoutes.GET("/:id")
	userRoutes.PUT("/")
	userRoutes.DELETE("/:id")
}
