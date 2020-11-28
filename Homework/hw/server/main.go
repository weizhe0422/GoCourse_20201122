package main

import (
	"github.com/gin-gonic/gin"
	PILI "github.com/weizhe0422/GoCourse_20201122/Homework/hw/PILI/Controller"
	"golang.org/x/net/context"
	"net/http"
)

var PILIUtil *PILI.PILI

func main() {
	context.Background()
	router := gin.Default()
	http.HandleFunc("ss", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w,r)
	})

	PILIUtil = PILI.New()

	router.GET("/role", Get)
	router.GET("/role/:id", GetOne)
	router.POST("/role", Post)
	router.PUT("/role/:id", Put)
	router.DELETE("/role/:id", Delete)
	router.Run(":8080")
}
