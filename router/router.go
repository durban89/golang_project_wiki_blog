package router

import (
	"net/http"

	"github.com/durban89/wiki/controllers"
)

// RouterMap 路由
type RouterMap struct {
	Path string
	Fn   func(http.ResponseWriter, *http.Request)
}

// RouterMaps 路由列表
var RouterMaps = []*RouterMap{
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
