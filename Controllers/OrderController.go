package Controllers

import (
	"one_test_case/Helpers"
	"one_test_case/Models"
	"time"

	"github.com/gin-gonic/gin"
)

func ListOrder(c *gin.Context) {
	var orders []Models.Order

	err := Models.GetAllOrders(&orders)
	if err != nil {

		Helpers.RespError(c, err.Error())
		return
	} else {

		Helpers.RespOK(c, orders)
	}
}

func SaveOrder(c *gin.Context) {
	var order Models.Order
	var requestBody Models.OrderSaveReqBody

	if err := c.BindJSON(&requestBody); err != nil {

		Helpers.RespError(c, err.Error())
		return
	} else {

		c.BindJSON(&order)

		var product Models.Product
		err = Models.GetProductByCode(&product, requestBody.Code)
		if err != nil {
			Helpers.RespError(c, err.Error())
			return
		}
		order.ProductId = int32(product.Id)
		order.Quantity = requestBody.Quantity
		order.CreatedDate = time.Now().Format("2006-01-02 15:04:05")
		if product.Stock < order.Quantity {
			Helpers.RespError(c, "Stoktaki üründen fazlası girilmiş.")
			return
		}
		err := Models.CreateOrder(&order)
		if err != nil {

			Helpers.RespError(c, err.Error())
			return
		} else {

			Helpers.RespOK(c, "İşlem Başarılı")
		}
	}
}

func GetOrder(c *gin.Context) {

	var order Models.Order
	var requestBody Models.GetReqBody

	if err := c.BindJSON(&requestBody); err != nil {
		Helpers.RespError(c, err.Error())
		return
	} else {

		err := Models.GetOrderById(&order, requestBody.Id)
		if err != nil {
			Helpers.RespError(c, err.Error())
		}

		Helpers.RespOK(c, order)

	}
}
