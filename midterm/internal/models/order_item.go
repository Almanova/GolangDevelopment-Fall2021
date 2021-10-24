package models

type OrderItem struct {
	Id int `json:"id"`
	OrderId int `json:"order_id"`
	ProductId int `json:"product_id"`
	Quantity int `json:"quantity"`
	Price int `json:"price"`
	Discount int `json:"discount"`
	TotalPrice float64 `json:"total_price"`
}
