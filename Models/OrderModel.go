package Models

type Order struct {
	Id          uint   `json:"id"`
	ProductId   int32  `json:"product_id"`
	Quantity    int32  `json:"quantity"`
	CreatedDate string `json:"created_date"`
	Status      int32  `json:"status"`
}

func (o *Order) TableName() string {
	return "order"
}
