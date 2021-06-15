package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-web-collection/config"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("html/*")

	r.GET("/", index)

	r.Run(config.GlobalConfig.Port)
}

// 返回静态页
func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
