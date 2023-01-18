package main

import (
	"fmt"
	"github.com/spf13/viper"
	"m800/db"
)

func main() {
	viper.SetDefault("mongo.url", "mongodb://localhost:27017")
	db.NewMongoImpl(viper.GetString("mongo.url"))
	fmt.Println("Hello, World!")
}
