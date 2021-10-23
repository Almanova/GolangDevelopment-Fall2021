package models

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Category string `json:"category"`
	Brand string `json:"brand"`
	Description string `json:"description"`
	Weight int `json:"weight"`
	Price int `json:"price"`
}
