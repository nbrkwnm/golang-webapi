package router

import "github.com/gin-gonic/gin"

func Init() {
	router := gin.Default()

	initializeRoutes(router)

	router.Run("localhost:8080")
}
