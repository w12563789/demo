package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.Default()
	app.GET("/index", func(this *gin.Context) {
		this.Writer.Write([]byte("helle gin"))
	})
	app.Run()
}
