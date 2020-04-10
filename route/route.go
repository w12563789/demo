package route

import (
	"database/sql"
	controller "demo/controller"
	"fmt"
	"github.com/gin-gonic/gin"
)
func InitRouter() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?parseTime=true")
	defer db.Close()
	if err != nil{
		fmt.Println(err.Error())
	}
	//使用gin的Default方法创建一个路由handler
	app := gin.Default()
	//app.LoadHTMLGlob(filepath.Join(os.Getenv("GOPATH"),"src/demo/view/**/*"))

	//使用以下gin提供的Group函数为不同的API进行分组
	login := app.Group("login")
	{
		login.POST("/", controller.Login)
	}
	//监听 服务器 端口
	app.Run(":8080")
}
