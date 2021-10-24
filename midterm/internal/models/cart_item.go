package models

type CartItem struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	ProductId int `json:"product_id"`
	Quantity int `json:"quantity"`
}

type CartItemFull struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	Product ProductFull `json:"product"`
	Quantity int `json:"quantity"`
}
