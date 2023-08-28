package Routes

import (
	"one_test_case/Controllers"

	"github.com/gin-gonic/gin"
)

func CampaignRoutes(route *gin.Engine) {

	campaign := route.Group("/campaign")
	{

		campaign.GET("list", Controllers.ListCampaigns)
		campaign.POST("save", Controllers.SaveCampaign)

	}

}
