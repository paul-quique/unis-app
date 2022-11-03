package main

import (
	"github.com/gin-gonic/gin"
	api "github.com/paul-quique/unis"
)

func main() {
	r := gin.Default()
	r.GET("/user/:id", api.GetUser)
	r.POST("/auth", api.Auth)
	r.GET("/category/:id", api.GetCategory)
	r.GET("/categories", api.GetCategories)
	r.POST("/user", api.PostUser)
	r.Run()
}
