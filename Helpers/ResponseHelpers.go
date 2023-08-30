package Helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RespOK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "İşlem Başarılı",
		"data":    data,
	})

	return
}

func RespError(c *gin.Context, msg string) {

	c.JSON(http.StatusOK, msg)

	return
}
