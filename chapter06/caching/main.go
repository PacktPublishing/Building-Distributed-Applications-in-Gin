package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func IllustrationHandler(c *gin.Context) {
	c.Header("Etag", "illustration")
	c.Header("Cache-Control", "max-age=2592000")

	if match := c.GetHeader("If-None-Match"); match != "" {
		if strings.Contains(match, "illustration") {
			c.Writer.WriteHeader(http.StatusNotModified)
			return
		}
	}

	c.File("illustration.png")
}

func main() {
	router := gin.Default()
	router.GET("/illustration", IllustrationHandler)
	router.Run(":3000")
}
