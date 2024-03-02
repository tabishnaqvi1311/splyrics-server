package controllers

import (
	"context"
	"example/lyrics-server-go/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


var lyricColl *mongo.Collection = config.GetCollection(config.DB)

func GetLyricsById() gin.HandlerFunc {
	return func(c *gin.Context){
		id := c.Param("id")

		var res bson.M
		
		err := lyricColl.FindOne(context.Background(), bson.M{"id": id}).Decode(&res)

		if err != nil{
			c.JSON(http.StatusNotFound, gin.H{"error": "lyrics not found"})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}