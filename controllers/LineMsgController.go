package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LineMsgController struct{}

func (c *LineMsgController) Query(g *gin.Context) {

	g.JSON(http.StatusOK, "Test")
}
