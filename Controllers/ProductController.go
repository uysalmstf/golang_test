package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListProducts(c *gin.Context) {
	c.JSON(http.StatusOK, "merhaba")
}
