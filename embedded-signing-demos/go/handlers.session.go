package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func showIndexPage(c *gin.Context) {
	response := createEmbeddedEnvelope(c)
	iFrameURL := getIframeURL(response.Recipients[0].Token)

	log.Printf("Rendering iFrameURL: %s", iFrameURL)
	render(c, gin.H{
		"iframeURL": iFrameURL},
		"index.html")
	return
}

func showSuccessPage(c *gin.Context) {
	render(
		c,
		gin.H{
			"title": "Success",
		},
		"success.html",
	)
	return
}
