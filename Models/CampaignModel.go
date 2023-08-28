package Models

type Campaign struct {
	Id                     uint    `json:"id"`
	Name                   string  `json:"name"`
	ProductId              int32   `json:"product_id"`
	Duration               int32   `json:"duration"`
	PriceManipulationLimit int32   `json:"price_manipulation_limit"`
	TargetSalesCount       int32   `json:"target_sales_count"`
	Status                 int32   `json:"status"`
	CreatedDate            string  `json:"created_date"`
	PriceNow               float32 `json:"price_now"`
}

func (c *Campaign) TableName() string {
	return "campaign"
}
