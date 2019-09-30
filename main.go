package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

var Content string

func main(){
	port := os.Getenv("PORT")
	if port == "" {port = "8080"}
	router := gin.Default()
	router.Static("/_nuxt","./client/dist/_nuxt")
	router.GET("/", func(c *gin.Context){
    	c.File("./client/dist/index.html")
	})

	api := router.Group("/api")
	{
	    api.GET("/get", get)
	    api.POST("/post", post)
    }
	err := router.Run(":" + port)
	if err != nil {panic(err)}
}

func get(c *gin.Context){
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, Canning{Content})
}

func post(c *gin.Context){
	Content = c.Request.FormValue("content")
	print(Content)
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, Canning{Content})
}

type Canning struct{
	Content string `json:"content"`
}