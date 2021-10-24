package models

type Order struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	Products []*ProductFull `json:"products"`
	Status string `json:"status"`
	Address string `json:"address"`
	TotalPrice float64 `json:"total_price"`
}

type OrderDto struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	Status string `json:"status"`
	Address string `json:"address"`
	TotalPrice float64 `json:"total_price"`
}
