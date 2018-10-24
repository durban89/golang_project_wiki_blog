package router

import (
	"github.com/durban89/wiki/controllers"
	"github.com/durban89/wiki/controllers/article"
	"github.com/durban89/wiki/controllers/welcome"
)

// Routes 操作
func Routes() {
	// 文章
	GET("/articles/update/", article.Update)
	GET("/articles/view/", article.View)
	GET("/articles/", article.Item)
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
