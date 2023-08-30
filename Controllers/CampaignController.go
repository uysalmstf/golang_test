package Controllers

import (
	"fmt"
	"math/rand"
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

			lastUpdateParsedTime, err := time.Parse(layout, campaigns[i].LastUpdateDate)
			if err != nil {
				fmt.Println("Error:", err)
				c.JSON(http.StatusBadRequest, err.Error())
				continue
			}

			lastTime := lastUpdateParsedTime.Add(time.Duration(campaigns[i].PriceDuration) * time.Hour)

			if curParsedTime.After(lastTime) {

				//todo: price update
				campaigns[i].PriceNow = calcPrice(campaigns[i])
				campaigns[i].LastUpdateDate = time.Now().Format("2006-01-02 15:04:05")
				upErr := Models.UpdateCampaign(&campaigns[i])
				if upErr != nil {
					fmt.Println("Error:", upErr)
					c.JSON(http.StatusBadRequest, upErr.Error())
				}
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

		randDuration := generateRandomInt(1, int(campaign.Duration))

		campaign.PriceDuration = int32(randDuration)

		campaign.LastUpdateDate = campaign.CreatedDate

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

func GetCampaign(c *gin.Context) {
	var campaign Models.Campaign
	var requestBody Models.GetCampaignReqBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusNotFound, err.Error())
	} else {

		err := Models.GetCampaignByName(&campaign, requestBody.Name)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
		}

		c.JSON(http.StatusOK, campaign)

	}
}

func calcPrice(campaign Models.Campaign) float32 {

	min := campaign.PriceNow - (campaign.PriceNow * float32(campaign.PriceManipulationLimit) / 100)
	max := campaign.PriceNow + (campaign.PriceNow * float32(campaign.PriceManipulationLimit) / 100)

	newPrice := generateFloatWithinRange(min, max)

	return newPrice
}

func generateFloatWithinRange(minVal, maxVal float32) float32 {
	return minVal + rand.Float32()*(maxVal-minVal)
}

func generateRandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}
