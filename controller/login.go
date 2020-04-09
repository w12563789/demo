package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(this *gin.Context)  {

	if this.Request.Method == "POST" {
		var msg struct {
			Name    string `json:"user"`
			Message string
			Age     int
		}
		msg.Name = "小王子"
		msg.Message = "Hello world!"
		msg.Age = 18
		this.JSON(http.StatusOK,msg)
	}
	this.HTML(http.StatusOK, "login/index.html",nil)
}
