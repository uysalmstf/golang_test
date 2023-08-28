package Controllers

import (
	"fmt"
	"net/http"
	"one_test_case/Models"
	"time"

	"github.com/gin-gonic/gin"
)

func ListCampaigns(c *gin.Context) {
	var campaigns []Models.Campaign
	var newCampaignsArr []Models.Campaign

	layout := "2006-01-02 15:04:05"

	err := Models.GetAllCampaigns(&campaigns)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {

		for i := 0; i < len(campaigns); i++ {
			fmt.Println(campaigns[i].CreatedDate)

			parsedTime, err := time.Parse(layout, campaigns[i].CreatedDate)
			if err != nil {
				fmt.Println("Error:", err)
				c.JSON(http.StatusBadRequest, err.Error())
				continue
			}
			curParsedTime, curErr := time.Parse(layout, time.Now().Format("2006-01-02 15:04:05"))
			if curErr != nil {
				fmt.Println("Error:", curErr)
				c.JSON(http.StatusBadRequest, curErr.Error())
				continue
			}

			end_date := parsedTime.Add(time.Duration(campaigns[i].Duration) * time.Hour)
			if end_date.Before(curParsedTime) {

				campaigns[i].Status = 0

				upErr := Models.UpdateCampaign(&campaigns[i])
				if upErr != nil {
					fmt.Println("Error:", upErr)
					c.JSON(http.StatusBadRequest, upErr.Error())
				}
				continue
			}

			newCampaignsArr = append(newCampaignsArr, campaigns[i])

		}
		c.JSON(http.StatusOK, newCampaignsArr)
	}
}

func SaveCampaign(c *gin.Context) {

	var product Models.Product
	var campaign Models.Campaign
	var requestBody Models.CampaignSaveReqBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusNotFound, err.Error())
	} else {

		err := Models.GetProductByCode(&product, requestBody.Code)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
		}

		campaign.Duration = requestBody.Duration
		campaign.Name = requestBody.Name
		campaign.PriceManipulationLimit = requestBody.PriceManipulationLimit
		campaign.TargetSalesCount = requestBody.TargetSalesCount
		campaign.ProductId = int32(product.Id)
		campaign.PriceNow = product.Price
		campaign.Status = 1
		campaign.CreatedDate = time.Now().Format("2006-01-02 15:04:05")

		insertErr := Models.CreateCampaign(&campaign)
		if insertErr != nil {
			fmt.Println(insertErr.Error())
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  true,
				"message": "İşlem Başarılı",
			})
		}
	}
}
