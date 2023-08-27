package Models

type Product struct {
	Id    uint    `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
	Code  string  `json:"code"`
	Stock int32   `json:"stock"`
}

func (p *Product) TableName() string {
	return "product"
}
