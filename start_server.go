package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func startServer(webPath ...string) {
	router := gin.Default()
	if len(webPath) != 0 {
		router.LoadHTMLGlob(webPath[0] + "/index.html")
		router.Static("/static", webPath[0]+"/static")
	}
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.Run(":7684")
}
