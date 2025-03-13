package Server

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/infra/http/paths"
)

func GetRouter() *gin.Engine{
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Content-Type"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}

	router.Use(cors.New(config))
	paths.AddUserRoutes(router)
	return router
}