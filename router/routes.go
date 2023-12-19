package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nbrkwnm/golang-webapi/handler"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/albums", handler.GetAlbumsHandler)
		v1.GET("/albums/:id", handler.GetAlbumByIdHandler)
		v1.POST("/albums", handler.CreateAlbumHandler)
		v1.PUT("/albums/:id", handler.UpdateAlbumHandler)
		v1.DELETE("/albums/:id", handler.DeleteAlbumHandler)
	}
}
