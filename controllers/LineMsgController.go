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

type LineMsgController struct {
	bot *linebot.Client
}

func NewLineMsgController() LineMsgController {
	// 設定 line bot client
	bot, err := linebot.New(
		viper.GetString("line.secret"),
		viper.GetString("line.token"),
	)
	if err != nil {
		log.Fatal(err)
	}
	return LineMsgController{
		bot,
	}
}

func (c *LineMsgController) Query(g *gin.Context) {

	g.JSON(http.StatusOK, "Test")
}

func (c *LineMsgController) Save(g *gin.Context) {
	// 取得請求中的資料
	events, err := c.bot.ParseRequest(g.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			g.JSON(400, gin.H{"message": "Invalid signature"})
		} else {
			g.JSON(500, gin.H{"message": "Internal Server Error"})
		}
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			// 將資料儲存到 MongoDB
			saveMessage(event)
		}
	}
	g.JSON(200, gin.H{"message": "OK"})
}

func (c *LineMsgController) Send(g *gin.Context) {
	// 從請求中取得 userID 和 message
	var req struct {
		UserID  string `json:"userID"`
		Message string `json:"message"`
	}
	if err := g.BindJSON(&req); err != nil {
		g.AbortWithStatus(http.StatusBadRequest)
		return
	}
	// 回傳訊息
	// 建立文字訊息
	msg := linebot.NewTextMessage(req.Message)
	// 回傳訊息
	_, err := c.bot.PushMessage(req.UserID, msg).Do()
	if err != nil {
		log.Print(err)
	}
	g.JSON(http.StatusOK, gin.H{})
}

func saveMessage(event *linebot.Event) {
	msg, _ := event.Message.(*linebot.TextMessage)
	// Create a message
	message := dto.Message{
		UserID: event.Source.UserID,
		Text:   msg.Text,
	}

	mongoImpl := db.NewMongoImpl()
	mongoImpl.Save(message)
}
