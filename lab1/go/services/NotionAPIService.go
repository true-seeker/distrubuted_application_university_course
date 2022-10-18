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

const AESKey = "the-key-has-to-be-32-bytes-long!"

func NewNotionAPI() (NotionAPI, error) {
	notionAPI := NotionAPI{}
	var notionCredentials dto.NotionCredentialsDTO
	a := AES{key: []byte(AESKey)}
	encryptedData, err := ReadFile()
	if err != nil {
		log.Fatal(err)
	}

	decrypted, err := a.Decrypt(encryptedData)
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

func (n *NotionAPI) DeletePageById(pageId string) error {
	deletePageByIdUrl := fmt.Sprintf("https://api.notion.com/v1/pages/%s", pageId)
	client := &http.Client{}

	values := map[string]bool{
		"archived": true,
	}

	jsonData, err := json.Marshal(values)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PATCH", deletePageByIdUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", n.accessToken))
	req.Header.Set("Notion-Version", "2022-06-28")
	req.Header.Set("Content-Type", "application/json")
	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func (n *NotionAPI) AddPage(properties map[string]interface{}, databaseId string) error {
	addPageUrl := fmt.Sprintf("https://api.notion.com/v1/pages/")
	client := &http.Client{}

	values := map[string]interface{}{
		"parent": map[string]string{
			"database_id": databaseId,
		},
		"properties": properties,
	}

	jsonData, err := json.Marshal(values)
	if err != nil {
		return err
	}
	fmt.Println(string(jsonData))

	req, err := http.NewRequest("POST", addPageUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", n.accessToken))
	req.Header.Set("Notion-Version", "2022-06-28")
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}

func (n *NotionAPI) UpdatePage(properties map[string]interface{}, pageId string) error {
	updatePageUrl := fmt.Sprintf("https://api.notion.com/v1/pages/%s", pageId)
	client := &http.Client{}

	values := map[string]interface{}{
		"properties": properties,
	}

	jsonData, err := json.Marshal(values)
	if err != nil {
		return err
	}
	fmt.Println(string(jsonData))

	req, err := http.NewRequest("PATCH", updatePageUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", n.accessToken))
	req.Header.Set("Notion-Version", "2022-06-28")
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}
