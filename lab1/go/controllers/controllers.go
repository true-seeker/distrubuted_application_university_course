package controllers

import (
	"log"
	"main/services"

	"github.com/gin-gonic/gin"
)

func Controllers() {
	r := gin.Default()
	err := r.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal(err)
	}

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", services.Index)
	r.GET("/notion_auth", services.NotionAuthRedirect)
	r.GET("/logout", services.Logout)
	r.Run(":80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
