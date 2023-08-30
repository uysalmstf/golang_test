package Controllers

import (
	"fmt"
	"net/http"
	"one_test_case/Models"

	"github.com/gin-gonic/gin"
)

func ListProducts(c *gin.Context) {
	var products []Models.Product

	err := Models.GetAllProducts(&products)
	if err != nil {

		c.AbortWithStatus(http.StatusNotFound)
	} else {

		c.JSON(http.StatusOK, products)
	}
}

func SaveProduct(c *gin.Context) {
	var product Models.Product
	var requestBody Models.ProductSaveReqBody

	if err := c.BindJSON(&requestBody); err != nil {

		c.JSON(http.StatusNotFound, err.Error())
	} else {

		c.BindJSON(&product)
		product.Name = requestBody.Name
		product.Code = requestBody.Code
		product.Stock = requestBody.Stock
		product.Price = requestBody.Price

		err := Models.CreateProduct(&product)
		if err != nil {

			fmt.Println(err.Error())
			c.AbortWithStatus(http.StatusNotFound)
		} else {

			c.JSON(http.StatusOK, gin.H{
				"status":  true,
				"message": "İşlem Başarılı",
			})
		}
	}
}

func GetProduct(c *gin.Context) {
	var product Models.Product
	var requestBody Models.GetProductReqBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusNotFound, err.Error())
	} else {

		err := Models.GetProductByCode(&product, requestBody.Code)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
		}

		c.JSON(http.StatusOK, product)

	}
}
