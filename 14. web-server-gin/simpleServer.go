package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// gin.Default just do the following two line
// r := gin.New()
// r.Use(gin.Logger(), gin.Recovery())

func simpleServer() {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, map[string]string{
			"message": "Hello world",
		})
	})

	r.Run(":8088")
}


