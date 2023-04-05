package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func startServer(webPath ...string) {
	gin.SetMode(gin.ReleaseMode) // Disable gin's stdout.
	router := gin.New()          // The release mode doesn't work with `gin.Default()`.

	if len(webPath) != 0 {
		router.LoadHTMLGlob(webPath[0] + "/index.html")
		router.Static("/static", webPath[0]+"/static")
		router.GET("/web", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})
		webpageRunningNotification()
	}

	router.Run(PORT)
}

func webpageRunningNotification() {
	fmt.Println("The LANWorker webpage is now running,")
	fmt.Printf("visit http://%s%s/web to use LANWorker-web. \n", IP, PORT)
	fmt.Println("Make sure your device is in the same LAN with LANWorker-desktop.")
}
