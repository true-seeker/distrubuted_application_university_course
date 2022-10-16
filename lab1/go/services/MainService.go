package services

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func Index(c *gin.Context) {
	if _, err := os.Stat("test.txt"); err == nil {

		api, err := NewNotionAPI()
		if err != nil {
			log.Fatal(err)
		}
		databases, err := api.FindDatabases()
		if err != nil {
			log.Fatal(err)
		}
		//for i := 0; i < len(databases.Results); i++ {
		//	fmt.Println(databases.Results[i])
		//	fmt.Println("+++++++++++")
		//}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":     "Lab1",
			"user":      api.user,
			"authUrl":   "https://api.notion.com/v1/oauth/authorize?client_id=710003c6-cbb2-4b1f-b979-248a38a1d2db&response_type=code&owner=user&redirect_uri=http%3A%2F%2Flocalhost%2Fnotion_auth",
			"databases": databases.Results,
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
	c.Redirect(http.StatusFound, "/")
}

func Database(c *gin.Context) {

}
