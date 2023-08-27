package Routes

import (
	"one_test_case/Controllers"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {

	product := route.Group("/product")
	{

		product.GET("list", Controllers.ListProducts)
		product.POST("save", Controllers.SaveProduct)

	}

}
