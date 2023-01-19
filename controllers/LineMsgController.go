package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/spf13/viper"
	"log"
	"m800/db"
	"m800/interal/dto"
	"net/http"
)

type LineMsgController struct{}

func (c *LineMsgController) Query(g *gin.Context) {

	g.JSON(http.StatusOK, "Test")
}

func (c *LineMsgController) Send(g *gin.Context) {

	//Call line
	//Save msg to
	// 設定 linebot client
	bot, err := linebot.New(
		viper.GetString("line.secret"),
		viper.GetString("line.token"),
	)
	if err != nil {
		log.Fatal(err)
	}
	// 發送文字訊息
	if _, err := bot.PushMessage("USER_ID", linebot.NewTextMessage("Hello World!")).Do(); err != nil {
		log.Print(err)
	}
	g.JSON(http.StatusOK, "Sent")
}

func (c *LineMsgController) Save(g *gin.Context) {
	lineWebhookHandler(g)
}

func lineWebhookHandler(c *gin.Context) {
	// 設定 linebot client
	bot, err := linebot.New(
		viper.GetString("line.secret"),
		viper.GetString("line.token"),
	)
	// 取得請求中的資料
	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.JSON(400, gin.H{"message": "Invalid signature"})
		} else {
			c.JSON(500, gin.H{"message": "Internal Server Error"})
		}
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			// 將資料儲存到 MongoDB
			saveMessageToMongoDB(event)
		}
	}
	c.JSON(200, gin.H{"message": "OK"})
}

func saveMessageToMongoDB(event *linebot.Event) {
	msg, _ := event.Message.(*linebot.TextMessage)
	// Create a message
	message := dto.Message{
		UserID: event.Source.UserID,
		Text:   msg.Text,
	}

	mongoImpl := db.NewMongoImpl()
	mongoImpl.Save(message)
}
