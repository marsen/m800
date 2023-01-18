package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"m800/controllers"
	"m800/db"
)

func main() {
	viper.SetDefault("mongo.url", "mongodb://localhost:27017")
	db.NewMongoImpl(viper.GetString("mongo.url"))
	g := gin.Default()
	g.Use(cors.Default())
	// 創建一個新的 GameController。
	gc := &controllers.LineMsgController{}
	game := g.Group("game")
	game.GET("query", gc.NewGame)
	g
	fmt.Println("Hello, World!")
}
