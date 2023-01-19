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
	fmt.Println("Hello, World!")
}
