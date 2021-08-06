package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

const (

	defaultBaseURL string = "https://www.yotisign.com"
)

func showIndexPage(c *gin.Context) {
	recipientToken  := getRecipientToken()

	if recipientToken == "" {
		c.HTML(
			http.StatusBadRequest,
			"error.html",
			gin.H{
				"ErrorTitle":   "Error when trying to retrieve recipient token",
				"ErrorMessage": "missing"})
		return
	}

	iFrameURL := getIframeURL(recipientToken)

	log.Printf("Rendering iFrameURL: %s", iFrameURL)
	render(c, gin.H{
		"iframeURL": iFrameURL},
		"index.html")
	return
}

// TODO: remove, get this from a new request
func getRecipientToken() string {
	if value, exists := os.LookupEnv("RECIPIENT_TOKEN"); exists && value != "" {
		return value
	}

	return ""
}

func getBaseURL() string {
	if value, exists := os.LookupEnv("YOTI_SIGN_BASE_URL"); exists && value != "" {
		return value
	}

	return defaultBaseURL
}

func getIframeURL(recipientToken string) string {
	baseURL := getBaseURL()
	return fmt.Sprintf("%s/embedded/sign/%s", baseURL, recipientToken)
}

func showSuccessPage(c *gin.Context) {
	render(
		c,
		gin.H{
			"title":            "Success",
		},
		"success.html",
	)
	return
}
