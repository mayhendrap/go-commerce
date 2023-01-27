package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	router.Use(cors.Default())
	return router
}
