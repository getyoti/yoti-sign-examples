package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)


const (
	selfSignedKeyName  = "SelfSignedKey.pem"
	selfSignedCertName = "SelfSignedCert.pem"
	portNumber         = 8080
)

func home(w http.ResponseWriter, r *http.Request) {
	token := "Bearer " + os.Getenv("YOTI_AUTHENTICATION_TOKEN")
	url := os.Getenv("YOTI_SIGN_BASE_URL")

	//Read options.json
	optionsFile,err := os.Open("options.json")
	defer optionsFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	options,err := ioutil.ReadAll(optionsFile)


	//build payload
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("options", string(options))

	//read test.pdf
	file, err := os.Open("test.pdf")
	defer file.Close()
	part2, err := writer.CreateFormFile("file",filepath.Base("test.pdf"))
	_, err = io.Copy(part2, file)
	err = writer.Close()

	req, err := http.NewRequest("POST", url, payload)
	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "multipart/form-data")
	req.Header.Set("Content-Type", writer.FormDataContentType())


	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		w.Write(data)
	}
}

func main() {
	_, insecureTLS := os.LookupEnv("TLS_INSECURE")
	if insecureTLS {
		fmt.Println("WARNING: TLS Certificate checking disabled")
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	certificatePresent := certificatePresenceCheck(selfSignedCertName, selfSignedKeyName)
	if !certificatePresent {
		fmt.Println("Generating self-signed certificates")
		err := generateSelfSignedCertificate(
			selfSignedCertName,
			selfSignedKeyName,
			fmt.Sprintf("127.0.0.1:%d", portNumber),
		)
		if err != nil {
			panic("Error creating certs: " + err.Error())
		}
	}

	http.HandleFunc("/", home)

	fmt.Printf("Starting server on port %d\n", portNumber)
	err := http.ListenAndServeTLS(fmt.Sprintf(":%d", portNumber), selfSignedCertName, selfSignedKeyName, nil)
	if err != nil {
		panic("Error when calling `ListenAndServeTLS`: " + err.Error())
	}
}
