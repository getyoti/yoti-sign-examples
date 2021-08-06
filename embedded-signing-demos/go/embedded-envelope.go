package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type EmbeddedEnvelopeResponse struct {
	EnvelopeID string      `json:"envelope_id"`
	Recipients []Recipient `json:"recipients"`
}

type Recipient struct {
	Token string `json:"token"`
	Email string `json:"email"`
}

func createEmbeddedEnvelope(c *gin.Context) (envResponse EmbeddedEnvelopeResponse) {
	// Read options.json
	optionsFile, err := os.Open("options.json")
	defer optionsFile.Close()
	if err != nil {
		setPreconditionError(c, "Error reading options.json", err.Error())
		return
	}

	options, err := ioutil.ReadAll(optionsFile)
	if err != nil {
		setPreconditionError(c, "Error getting options bytes", err.Error())
		return
	}
	// build payload
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	err = writer.WriteField("options", string(options))
	if err != nil {
		setPreconditionError(c, "Error getting options bytes", err.Error())
		return
	}

	// read test.pdf
	file, err := os.Open("test.pdf")
	defer file.Close()
	if err != nil {
		setPreconditionError(c, "Error getting test PDF", err.Error())
		return
	}

	part2, err := writer.CreateFormFile("file", filepath.Base("test.pdf"))
	if err != nil {
		setPreconditionError(c, "Error creating form", err.Error())
		return
	}

	_, err = io.Copy(part2, file)
	if err != nil {
		setPreconditionError(c, "Error copying file", err.Error())
		return
	}

	err = writer.Close()
	if err != nil {
		setPreconditionError(c, "Error closing writer", err.Error())
		return
	}

	embeddedEnvelopeURL := getEmbeddedEnvelopeURL()
	envBearerToken := os.Getenv("YOTI_AUTHENTICATION_TOKEN")

	if envBearerToken == "" {
		setPreconditionError(c, "YOTI_AUTHENTICATION_TOKEN missing", "add this value to the .env file")
		return
	}

	bearerToken := "Bearer " + envBearerToken
	req, err := http.NewRequest("POST", embeddedEnvelopeURL, payload)
	req.Header.Add("Authorization", bearerToken)
	req.Header.Add("Content-Type", "multipart/form-data")
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		setInternalServerError(c, "Error making HTTP request", err.Error())
		return
	}
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		setInternalServerError(c, "Error reading response bytes", err.Error())
		return
	}

	if response.StatusCode != 202 {
		c.HTML(
			response.StatusCode,
			"error.html",
			gin.H{
				"ErrorTitle": "Error Response",
				"ErrorMessage": fmt.Errorf(
					"status code: %v, message: %v",
					response.StatusCode,
					string(responseBytes))})
		return
	}

	var result EmbeddedEnvelopeResponse
	err = json.Unmarshal(responseBytes, &result)
	if err != nil {
		setInternalServerError(c, "Error unmarshalling response", err.Error())
		return
	}

	return result
}
