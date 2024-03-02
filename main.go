package main

import (
	"example/lyrics-server-go/config"
	"example/lyrics-server-go/routes"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	config.ConnectToMongo()

	router.GET("/", func(g *gin.Context){
		g.JSON(200, gin.H{"test": "test data"})
	})

	routes.LyricRoute(router)

	router.Run("localhost:8181")
}