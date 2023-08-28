package Models

type CampaignSaveReqBody struct {
	Name                   string `json:"name"`
	Code                   string `json:"code"`
	Duration               int32  `json:"duration"`
	PriceManipulationLimit int32  `json:"price_manipulation_limit"`
	TargetSalesCount       int32  `json:"target_sales_count"`
}
