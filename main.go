package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"m800/controllers"
)

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.ReadInConfig()
	viper.SetDefault("mongo.url", "mongodb://localhost:27017")
	//db.NewMongoImpl(viper.GetString("mongo.url"))
	g := gin.Default()
	g.Use(cors.Default())
	// 創建一個新的 GameController。
	//gc := &controllers.LineMsgController{}
	gc := controllers.NewLineMsgController()
	game := g.Group("lineMsg")
	game.POST("send", gc.Send)
	game.POST("save", gc.Save)
	game.GET("query/:userId", gc.Query)
	fmt.Println("Hello, World!")
	g.Run()
}
