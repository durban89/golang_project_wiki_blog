package main

import (
	"log"
	"net/http"

	"github.com/durban89/wiki/controllers"
	"github.com/durban89/wiki/helpers"
)

func main() {
	http.HandleFunc("/view/", helpers.MakeHandler(controllers.ArticleView))
	http.HandleFunc("/save/", helpers.MakeHandler(controllers.ArticleSave))
	http.HandleFunc("/edit/", helpers.MakeHandler(controllers.ArticleEdit))
	http.HandleFunc("/upload/", controllers.UploadHandler)
	http.HandleFunc("/postFile/", controllers.PostFileHandler)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
