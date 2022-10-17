package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"main/dto"
	"net/http"
)

type NotionAPI struct {
	accessToken string
	user        string
}

func NewNotionAPI() (NotionAPI, error) {
	notionAPI := NotionAPI{}
	var notionCredentials dto.NotionCredentialsDTO
	a := AES{key: []byte("the-key-has-to-be-32-bytes-long!")}
	encrypted_data, err := ReadFile()
	if err != nil {
		log.Fatal(err)
	}

	decrypted, err := a.Decrypt(encrypted_data)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(decrypted, &notionCredentials)
	if err != nil {
		log.Fatal(err)
	}

	notionAPI.accessToken = notionCredentials.AccessToken
	notionAPI.user = notionCredentials.Owner.User.Name
	return notionAPI, nil
}

func (n *NotionAPI) FindDatabases() (dto.NotionSearchDTO, error) {
	findDatabaseUrl := "https://api.notion.com/v1/search"
	searchDTO := dto.NotionSearchDTO{}
	client := &http.Client{}

	values := map[string]map[string]string{
		"filter": {
			"value":    "database",
			"property": "object",
		},
		"sort": {
			"direction": "ascending",
			"timestamp": "last_edited_time",
		},
	}

	jsonData, err := json.Marshal(values)
	if err != nil {
		return dto.NotionSearchDTO{}, err
	}

	req, err := http.NewRequest("POST", findDatabaseUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return dto.NotionSearchDTO{}, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", n.accessToken))
	req.Header.Set("Notion-Version", "2022-06-28")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return dto.NotionSearchDTO{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return dto.NotionSearchDTO{}, err
	}
	err = json.Unmarshal(body, &searchDTO)

	return searchDTO, nil
}

func (n *NotionAPI) GetDatabaseById(id string) (dto.NotionSearchDTO, error) {
	getDatabaseByIdUrl := fmt.Sprintf("https://api.notion.com/v1/databases/%s/query", id)
	searchDTO := dto.NotionSearchDTO{}
	client := &http.Client{}

	req, err := http.NewRequest("POST", getDatabaseByIdUrl, nil)
	if err != nil {
		return dto.NotionSearchDTO{}, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", n.accessToken))
	req.Header.Set("Notion-Version", "2022-06-28")
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return dto.NotionSearchDTO{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return dto.NotionSearchDTO{}, err
	}
	err = json.Unmarshal(body, &searchDTO)

	//fmt.Println(searchDTO)
	return searchDTO, nil
}
