package routes

import (
	"example/lyrics-server-go/controllers"

	"github.com/gin-gonic/gin"
)

func LyricRoute(router *gin.Engine){
	router.GET("/api/lyrics/:id", controllers.GetLyricsById())
}