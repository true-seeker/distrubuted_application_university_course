package services

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

var AuthUrl = fmt.Sprintf("https://api.notion.com/v1/oauth/authorize?client_id=%s&response_type=code&owner=user&redirect_uri=%s", clientId, redirectUrl)
var CredentialsFileName = "test.txt"

func Index(c *gin.Context) {
	if _, err := os.Stat(CredentialsFileName); err == nil {

		api, err := NewNotionAPI()
		if err != nil {
			log.Fatal(err)
		}
		databases, err := api.FindDatabases()
		if err != nil {
			log.Fatal(err)
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":     "Lab1",
			"user":      api.user,
			"databases": databases.Results,
		})

	} else if errors.Is(err, os.ErrNotExist) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "Lab1",
			"authUrl": AuthUrl,
		})
	}
}

func NotionAuthRedirect(c *gin.Context) {
	code := c.Query("code")
	credentials := GetNotionCredentials(code)

	a := AES{key: []byte("the-key-has-to-be-32-bytes-long!")}
	encrypted, err := a.Encrypt(credentials)
	if err != nil {
		log.Fatal(err)
	}

	err = WriteFile(encrypted)
	if err != nil {
		log.Fatal(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/")
}

func Logout(c *gin.Context) {
	err := DeleteFile()
	if err != nil {
		log.Fatal(err)
	}
	c.Redirect(http.StatusFound, "/")
}

func Database(c *gin.Context) {
	id := c.Param("id")

	api, err := NewNotionAPI()
	if err != nil {
		log.Fatal(err)
	}

	searchDto, err := api.GetDatabaseById(id)
	if err != nil {
		log.Fatal(err)
	}

	fields := map[string]interface{}{}

	for _, result := range searchDto.Results {
		for key, value := range result.Properties.(map[string]interface{}) {
			fields[key] = value
		}
	}

	if _, err := os.Stat(CredentialsFileName); err == nil {

		c.HTML(http.StatusOK, "database.html", gin.H{
			"title":      "Lab1.Database",
			"user":       api.user,
			"results":    searchDto.Results,
			"fields":     fields,
			"databaseId": id,
		})

	} else if errors.Is(err, os.ErrNotExist) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "Lab1.Database",
			"authUrl": AuthUrl,
		})
	}
}

func PageDelete(c *gin.Context) {
	id := c.Param("id")
	databaseId := c.Request.FormValue("databaseId")

	fmt.Println(databaseId)
	api, err := NewNotionAPI()
	if err != nil {
		log.Fatal(err)
	}

	err = api.DeletePageById(id)
	if err != nil {
		log.Fatal(err)
	}
	c.Redirect(http.StatusFound, fmt.Sprintf("/database/%s", databaseId))
}

func PageAdd(c *gin.Context) {
	databaseId := c.Request.FormValue("databaseId")
	delete(c.Request.PostForm, "databaseId")
	properties := parsePage(c.Request.PostForm)

	api, err := NewNotionAPI()
	if err != nil {
		log.Fatal(err)
	}

	err = api.AddPage(properties, databaseId)
	if err != nil {
		log.Fatal(err)
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/database/%s", databaseId))
}

func PageUpdate(c *gin.Context) {
	id := c.Param("id")
	databaseId := c.Request.FormValue("databaseId")
	delete(c.Request.PostForm, "databaseId")
	properties := parsePage(c.Request.PostForm)

	api, err := NewNotionAPI()
	if err != nil {
		log.Fatal(err)
	}

	err = api.UpdatePage(properties, id)
	if err != nil {
		log.Fatal(err)
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/database/%s", databaseId))

}

func parsePage(r url.Values) map[string]interface{} {
	properties := map[string]interface{}{}
	for key, value := range r {
		fieldName := strings.Split(key, "___")[0]
		fieldType := strings.Split(key, "___")[1]
		fmt.Println(fieldName, fieldType, value)
		if fieldType == "number" && value[0] != "" {
			number, _ := strconv.Atoi(value[0])
			properties[fieldName] = map[string]interface{}{
				"number": number,
			}
		} else if fieldType == "checkbox" {
			properties[fieldName] = map[string]interface{}{
				"checkbox": true,
			}
		} else if fieldType == "title" {
			properties[fieldName] = map[string][]map[string]interface{}{
				"title": {{"type": "text",
					"text": map[string]interface{}{
						"content": value[0],
					},
				},
				},
			}
		} else if fieldType == "date" {
			properties[fieldName] = map[string]interface{}{
				"date": map[string]string{
					"start": value[0],
				},
			}
		}
	}
	return properties
}
