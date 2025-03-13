package main

import (
	Server "github.com/LuisGerardoDC/Orbi/UserService/src/internal/infra/http"
	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.DebugMode)
	router := Server.GetRouter()
	router.Run(":8080")

}
