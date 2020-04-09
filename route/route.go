package route

import (
	controller "demo/controller"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
)
func InitRouter() {
	//使用gin的Default方法创建一个路由handler
	app := gin.Default()
	app.LoadHTMLGlob(filepath.Join(os.Getenv("GOPATH"),"src/demo/view/**/*"))
	//使用以下gin提供的Group函数为不同的API进行分组
	login := app.Group("login")
	{
		login.GET("/", controller.Login)
	}
	//监听 服务器 端口
	app.Run(":8080")
}
