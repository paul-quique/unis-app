package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	api "github.com/paul-quique/unis"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/user/:id", api.GetUser)
	r.POST("/auth", api.Auth)
	r.GET("/category/:id", api.GetCategory)
	r.GET("/categories", api.GetCategories)
	r.POST("/offers", api.GetOffers)
	r.POST("/user", api.PostUser)
	r.POST("/product", api.PostProduct)
	r.POST("/product/:id", api.DeleteProduct)
	r.GET("/product/:id", api.GetProduct)
	r.POST("/offers/accepted", api.GetAcceptedOffers)
	r.POST("/offer", api.PostOffer)
	r.POST("/offer/accept", api.AcceptOffer)
	r.POST("/message", api.PostMessage)
	r.POST("/conversation", api.PostConversation)
	r.POST("/conversations", api.PostConversations)
	r.POST("/search", api.PostSearch)
	r.Run()
}
