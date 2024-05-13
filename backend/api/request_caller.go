package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var sessionToken string

func init() {
	requestNewSessionToken()
}

func getSessionRequest(url string) []byte {
	if sessionToken == "" {
		requestNewSessionToken()
	}
	sessionUrl := url + "&token=" + sessionToken
	return getRequest(sessionUrl)
}

func getRequest(url string) []byte {

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	return body
}

type TokenResponse struct {
	ResponseCode    int    `json:"response_code"`
	ResponseMessage string `json:"response_message"`
	Token           string `json:"token"`
}

func requestNewSessionToken() {

	var tokenResponse TokenResponse
	requestURL := "https://opentdb.com/api_token.php?command=request"

	res, err := http.Get(requestURL)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &tokenResponse)

	if err != nil {
		log.Fatal(err)
	}

	sessionToken = tokenResponse.Token
	fmt.Println("Session token has been updated!")
}
