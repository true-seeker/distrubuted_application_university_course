package services

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"main/dto"
	"net/http"
	"os"
)

func Index(c *gin.Context) {
	if _, err := os.Stat("test.txt"); err == nil {
		a := AES{key: []byte("the-key-has-to-be-32-bytes-long!")}
		encrypted_data, err := ReadFile()
		if err != nil {
			log.Fatal(err)
		}

		decrypted, err := a.Decrypt(encrypted_data)
		if err != nil {
			log.Fatal(err)
		}
		var notionCredentials dto.NotionCredentialsDTO
		err = json.Unmarshal(decrypted, &notionCredentials)
		if err != nil {
			log.Fatal(err)
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "Lab1",
			"user":    notionCredentials.Owner.User.Name,
			"authUrl": "https://api.notion.com/v1/oauth/authorize?client_id=710003c6-cbb2-4b1f-b979-248a38a1d2db&response_type=code&owner=user&redirect_uri=http%3A%2F%2Flocalhost%2Fnotion_auth",
			"code":    string(decrypted),
		})

	} else if errors.Is(err, os.ErrNotExist) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "Lab1",
			"authUrl": "https://api.notion.com/v1/oauth/authorize?client_id=710003c6-cbb2-4b1f-b979-248a38a1d2db&response_type=code&owner=user&redirect_uri=http%3A%2F%2Flocalhost%2Fnotion_auth",
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
	c.Redirect(http.StatusMovedPermanently, "/")
}
