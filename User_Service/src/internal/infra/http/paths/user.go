package paths

import (
	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/infra/http/handlers"
	"github.com/gin-gonic/gin"
)

func AddUserRoutes(router *gin.Engine) {
	userRoutes := router.Group("/user")
	userRoutes.POST("/", handlers.ImpNewUserHandler.Handle)
	userRoutes.GET("/:id,", handlers.ImpGetUserHandler.Handle)
	userRoutes.PUT("/", handlers.ImpUpdateUserHandler.Handle)
	userRoutes.DELETE("/:id", handlers.ImpDeleteUserHandler.Handle)
}
