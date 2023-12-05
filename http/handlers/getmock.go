package handlers

import (
	"net/http"
	"strconv"

	"farukh.go/micro/http/models"
	"github.com/gin-gonic/gin"
)

func GetMock(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.MockAlbum)
}

func PostMock(c *gin.Context) {
	var postData models.TestStruct
	c.BindJSON(&postData)
	models.MockAlbum = append(models.MockAlbum, postData)
	c.IndentedJSON(http.StatusOK, models.MockAlbum)
}

func GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if len(models.MockAlbum) > id {
		c.IndentedJSON(http.StatusOK, models.MockAlbum[id])
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no such Mock"})
	}
}
