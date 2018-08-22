package controllers

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/durban89/wiki/helpers"
)

// UploadHandler 上传文件
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)

	if r.Method == "GET" {
		p := &helpers.Page{Title: "Upload File"}

		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		p.Token = token
		helpers.RenderTemplate(w, "upload", p)
	} else {
		fmt.Println("Uploading ....")
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()

		fmt.Fprintf(w, "%v", handler.Header)

		fmt.Println("handler.Filename", handler.Filename)
		time := fmt.Sprintf("%d", time.Now().Unix())
		targetFileName := "data/" + time + handler.Filename

		f, err := os.OpenFile(targetFileName, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
		}

		defer f.Close()

		// 写入到文件中
		io.Copy(f, file)
	}

}

// PostFileHandler 模拟客户端表单提交
func PostFileHandler(w http.ResponseWriter, r *http.Request) {

	filename := r.URL.Query().Get("filename")
	targetURL := r.URL.Query().Get("targetURL")
	// filename string, targetURL string

	fmt.Println("filename = ", filename)
	fmt.Println("targetUrl = ", targetURL)

	sourceFileName := "data/" + filename

	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)

	// 这一步非常重要
	fileWriter, err := writer.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error writing to writer")
		return
	}

	// open file

	f, err := os.OpenFile(sourceFileName, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("error open file")
		return
	}
	defer f.Close()

	// io copy
	_, err = io.Copy(fileWriter, f)
	if err != nil {
		fmt.Println("error copy file")
		fmt.Println(err)
		return
	}

	contentType := writer.FormDataContentType()
	writer.Close()

	resp, err := http.Post(targetURL, contentType, buf)
	if err != nil {
		fmt.Println("error post file")
		return
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("error read file")
		return
	}

	fmt.Println(resp.Status)
	fmt.Println(string(respBody))

	return
}
