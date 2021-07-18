package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-web-collection/config"
	"github.com/golang-web-collection/utils"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("views/html/*")
	r.StaticFS("/static", http.Dir("./views/static"))

	r.GET("/", index)
	r.GET("/upload", uploadPage)
	r.POST("/upload", uploadPost)

	r.Run(config.ServerPort())
}

// 返回静态页
func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func uploadPage(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", nil)
}

func uploadPost(c *gin.Context) {
	row, _ := c.GetRawData()
	data := strings.Split(string(row), ",")
	suffix := data[0][strings.Index(data[0], "/")+1 : strings.Index(data[0], ";")]
	dist, _ := base64.StdEncoding.DecodeString(data[1])
	file := utils.FullFileName(suffix)
	fmt.Println(file)
	f, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()
	f.Write(dist)
	c.String(200, "error")
}
