package Routes

import (
	"one_test_case/Controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(route *gin.Engine) {

	campaign := route.Group("/order")
	{

		campaign.GET("list", Controllers.ListOrder)
		campaign.POST("save", Controllers.SaveOrder)
		campaign.POST("get", Controllers.GetOrder)

	}

}
