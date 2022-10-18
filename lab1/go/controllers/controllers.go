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
	r.GET("/database/:id", services.Database)
	r.POST("/page/:id/delete", services.PageDelete)
	r.Run(":8080")
}
