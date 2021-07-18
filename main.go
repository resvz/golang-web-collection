package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-web-collection/config"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("html/*")
	r.StaticFS("/static", http.Dir("./static"))

	r.GET("/", index)
	r.GET("/upload", upload)

	r.Run(config.GlobalConfig.Port)
}

// 返回静态页
func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func upload(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", nil)
}
