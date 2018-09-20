package router

import (
	"net/http"

	"github.com/durban89/wiki/controllers"
)

// RouterMap 路由
type Map struct {
	Path string
	Fn   func(http.ResponseWriter, *http.Request)
}

// RouterMaps 路由列表
var RouterMaps = []*Map{
	{
		Path: "/json/to/test",
		Fn:   controllers.JsonToTest,
	},
	{
		Path: "/json/process/",
		Fn:   controllers.JsonProcess,
	},
	{
		Path: "/json/to/interface",
		Fn:   controllers.JsonToInterface,
	},
	{
		Path: "/json/",
		Fn:   controllers.Json,
	},
	{
		Path: "/process/xml/",
		Fn:   controllers.WelcomeProcessXML,
	},
	{
		Path: "/xml/",
		Fn:   controllers.WelcomeXML,
	},
	{
		Path: "/login/",
		Fn:   controllers.WelcomeLogin,
	},
	{
		Path: "/item/",
		Fn:   controllers.ArticleItem,
	},
	{
		Path: "/view/",
		Fn:   controllers.ArticleItem,
	},
	{
		Path: "/save/",
		Fn:   controllers.ArticleViewWithID,
	},
	{
		Path: "/edit/",
		Fn:   controllers.ArticleEdit,
	},
	{
		Path: "/upload/",
		Fn:   controllers.UploadHandler,
	},
	{
		Path: "/postFile/",
		Fn:   controllers.PostFileHandler,
	},
}

// Routes 操作
func Routes() {
	for i := 0; i < len(RouterMaps); i++ {
		cRoute := RouterMaps[i]
		http.HandleFunc(cRoute.Path, cRoute.Fn)
	}
}
