package main

import (
	hr "farukh.go/micro/http/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Any("/mock", hr.GetMock)
	router.POST("/add-mock", hr.PostMock)
	router.GET("/:id", hr.GetById)
	router.Run("localhost:8080")
}
