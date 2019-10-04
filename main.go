package main

import (
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	"net/http"
	"os"
	"unsafe"
)

var Content string

func main(){
	port := os.Getenv("PORT")
	if port == "" {port = "8080"}
	router := gin.Default()

	ws := melody.New()

	ws.HandleMessage(func(m *melody.Session, msg[]byte){
		Content = *(*string)(unsafe.Pointer(&msg))
		err := ws.Broadcast(msg)
		if err != nil {panic(err)}
	})

	router.Static("/_nuxt","./client/dist/_nuxt")
	router.GET("/", func(c *gin.Context){
    	c.File("./client/dist/index.html")
	})

	router.GET("/ws", func(c *gin.Context){
		err := ws.HandleRequest(c.Writer, c.Request)
		if err != nil {panic(err)}
	})

	api := router.Group("/api")
	{
	    api.GET("/get", get)
    }
	err := router.Run(":" + port)
	if err != nil {panic(err)}
}

func get(c *gin.Context){
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, Canning{Content})
}

type Canning struct{
	Content string `json:"content"`
}
