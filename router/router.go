package router

import (
	"github.com/durban89/wiki/controllers"
	"github.com/durban89/wiki/controllers/article"
	"github.com/durban89/wiki/controllers/welcome"
)

// Routes 操作
//
// TODO
// 问题1： 不能使用同一个路由地址，会出现冲突，建议定义自己的路由方式
//
func Routes() {
	// 添加路由配置
	// 文章
	GET("/articles/update/", article.Update)
	GET("/articles/view/", article.View)
	GET("/articles/", article.Item)
	POST("/articles/save/", article.Save)

	// 其他
	GET("/process/xml/", welcome.WelcomeProcessXML)
	GET("/xml/", welcome.WelcomeXML)
	GET("/login/", welcome.WelcomeLogin)
	GET("/json/to/test/", controllers.JsonToTest)
	GET("/json/process/", controllers.JsonProcess)
	GET("/json/to/interface/", controllers.JsonToInterface)
	GET("/json/", controllers.Json)
	GET("/upload/", controllers.UploadHandler)
	GET("/postFile/", controllers.PostFileHandler)
}
