package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setPreconditionError(c *gin.Context, errorTitle string, err string) {
	c.HTML(
		http.StatusInternalServerError,
		"error.html",
		gin.H{
			"ErrorTitle":   errorTitle,
			"ErrorMessage": err})
	return
}

func setInternalServerError(c *gin.Context, errorTitle string, err string) {
	c.HTML(
		http.StatusBadRequest,
		"error.html",
		gin.H{
			"ErrorTitle":   errorTitle,
			"ErrorMessage": err})
	return
}
