package services

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var redirectUrl = "http://localhost:8080/notion_auth"
var tokenUrl = "https://api.notion.com/v1/oauth/token"
var clientId = "710003c6-cbb2-4b1f-b979-248a38a1d2db"
var clientSecret = "secret_OFmEcyjLzlXPUHwgQklloBT2TeT64fVkEjIJFVQ91Pk"

func GetNotionCredentials(code string) (response string) {
	credentials := fmt.Sprintf("%s:%s", clientId, clientSecret)
	b64Credentials := base64.StdEncoding.EncodeToString([]byte(credentials))
	client := &http.Client{}

	values := map[string]string{
		"code":         code,
		"grant_type":   "authorization_code",
		"redirect_uri": redirectUrl,
	}

	jsonData, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", tokenUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", b64Credentials))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	response = string(body)
	return
}
