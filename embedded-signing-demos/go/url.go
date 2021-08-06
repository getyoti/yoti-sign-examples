package main

import (
	"fmt"
	"os"
)

const (
	defaultBaseURL string = "https://demo.www.yotisign.com"
	defaultAPIURL  string = "https://demo.api.yotisign.com"
)

func getBaseURL() string {
	if value, exists := os.LookupEnv("YOTI_SIGN_BASE_URL"); exists && value != "" {
		return value
	}

	return defaultBaseURL
}

func getAPIURL() string {
	if value, exists := os.LookupEnv("YOTI_SIGN_API_URL"); exists && value != "" {
		return value
	}

	return defaultAPIURL
}

func getIframeURL(recipientToken string) string {
	baseURL := getBaseURL()
	return fmt.Sprintf("%s/embedded/sign/%s", baseURL, recipientToken)
}

func getEmbeddedEnvelopeURL() string {
	baseURL := getAPIURL()
	return fmt.Sprintf("%s/v2/embedded-envelopes", baseURL)
}
