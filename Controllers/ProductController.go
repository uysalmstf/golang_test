package Controllers

import (
	"one_test_case/Helpers"
	"one_test_case/Models"

	"github.com/gin-gonic/gin"
)

func ListProducts(c *gin.Context) {
	var products []Models.Product

	err := Models.GetAllProducts(&products)
	if err != nil {

		Helpers.RespError(c, err.Error())
		return
	} else {

		Helpers.RespOK(c, products)
	}
}

func SaveProduct(c *gin.Context) {
	var product Models.Product
	var requestBody Models.ProductSaveReqBody

	if err := c.BindJSON(&requestBody); err != nil {

		Helpers.RespError(c, err.Error())
		return
	} else {

		c.BindJSON(&product)
		product.Name = requestBody.Name
		product.Code = requestBody.Code

		var sameProduct Models.Product
		_ = Models.GetProductByCode(&sameProduct, product.Code)

		if sameProduct.Code != "" {
			Helpers.RespError(c, "Aynı kodda ürün mevcut.")
			return
		}

		product.Stock = requestBody.Stock
		product.Price = requestBody.Price

		err := Models.CreateProduct(&product)
		if err != nil {

			Helpers.RespError(c, err.Error())
			return
		} else {

			Helpers.RespOK(c, "İşlem Başarılı")
		}
	}
}

func GetProduct(c *gin.Context) {
	var product Models.Product
	var requestBody Models.GetProductReqBody

	if err := c.BindJSON(&requestBody); err != nil {
		Helpers.RespError(c, err.Error())
		return
	} else {

		err := Models.GetProductByCode(&product, requestBody.Code)
		if err != nil {
			Helpers.RespError(c, err.Error())
			return
		}

		Helpers.RespOK(c, product)

	}
}
