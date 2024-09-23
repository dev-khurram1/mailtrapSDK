package mailtrapSDK

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var internalToken string
var senderEmail string
var senderName string
var receiverEmail string
var receiverName string

var internalSubject string
var internalTextBody string

var testString = ""

var apiUrl = "https://send.api.mailtrap.io/api/send"

func SetToken(token string) {
	internalToken = token
}

func SetSender(email string, name string) {
	senderEmail = email
	senderName = name
}

func SetReceiver(email string, name string) {
	receiverEmail = email
	receiverName = name
}

func SetSubject(subject string) {
	internalSubject = subject
}

func SetBody(textBody string) {
	internalTextBody = textBody
}

func SendEmail() {
	jsonBody := map[string]interface{}{
		"to": []map[string]string{
			{"email": receiverEmail, "name": receiverEmail},
		},
		"from": map[string]string{
			"email": senderEmail, "name": senderName,
		},
		"custom_variables": map[string]string{
			"user_id":  "45982",
			"batch_id": "PSJ-12",
		},
		"headers": map[string]string{
			"X-Message-Source": "dev.mydomain.com",
		},
		"subject":  internalSubject,
		"text":     internalTextBody,
		"category": "API Test",
	}

	requestBody, err := json.Marshal(jsonBody)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// create new http request
	request, error := http.NewRequest("POST", apiUrl, bytes.NewBuffer(requestBody))
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Api-Token", internalToken)
	request.Header.Set("Content-Type", "application/json")

	// send the request
	client := &http.Client{}
	response, error := client.Do(request)

	if error != nil {
		fmt.Println(error)
	}

	responseBody, error := io.ReadAll(response.Body)

	if error != nil {
		fmt.Println(error)
	}

	formattedData := formatJSON(responseBody)
	fmt.Println("Status: ", response.Status)
	fmt.Println("Response body: ", formattedData)

	// clean up memory after execution
	defer response.Body.Close()
}

func formatJSON(data []byte) string {
	var out bytes.Buffer
	err := json.Indent(&out, data, "", "  ")

	if err != nil {
		fmt.Println(err)
	}

	d := out.Bytes()
	return string(d)
}
