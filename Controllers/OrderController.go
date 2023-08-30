package Controllers

import (
	"fmt"
	"net/http"
	"one_test_case/Models"
	"time"

	"github.com/gin-gonic/gin"
)

func ListOrder(c *gin.Context) {
	var orders []Models.Order

	err := Models.GetAllOrders(&orders)
	if err != nil {

		c.AbortWithStatus(http.StatusNotFound)
	} else {

		c.JSON(http.StatusOK, orders)
	}
}

func SaveOrder(c *gin.Context) {
	var order Models.Order
	var requestBody Models.OrderSaveReqBody

	if err := c.BindJSON(&requestBody); err != nil {

		c.JSON(http.StatusNotFound, err.Error())
	} else {

		c.BindJSON(&order)

		var product Models.Product
		err = Models.GetProductByCode(&product, requestBody.Code)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
		}
		order.ProductId = int32(product.Id)
		order.Quantity = requestBody.Quantity
		order.CreatedDate = time.Now().Format("2006-01-02 15:04:05")
		if product.Stock < order.Quantity {
			fmt.Println("stoktaki üründen fazlası girilmiş")
			c.AbortWithStatus(http.StatusNotFound)
		}
		err := Models.CreateOrder(&order)
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

func GetOrder(c *gin.Context) {

	var order Models.Order
	var requestBody Models.GetReqBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusNotFound, err.Error())
	} else {

		err := Models.GetOrderById(&order, requestBody.Id)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
		}

		c.JSON(http.StatusOK, order)

	}
}
