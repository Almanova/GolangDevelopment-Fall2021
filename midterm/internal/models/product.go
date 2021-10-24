package models

type Product struct {
	Id int `json:"id"`
	Name string `json:"name"`
	CategoryName string `json:"category_name"`
	BrandName string `json:"brand_name"`
	Description string `json:"description"`
	Weight int `json:"weight"`
	Price int `json:"price"`
	Discount int `json:"discount"`
}

type ProductFull struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Category Category `json:"category"`
	Brand Brand `json:"brand"`
	Description string `json:"description"`
	Weight int `json:"weight"`
	Price int `json:"price"`
	Discount int `json:"discount"`
}

type ProductDto struct {
	Id int `json:"id"`
	Name string `json:"name"`
	CategoryId int `json:"category_id"`
	BrandId int `json:"brand_id"`
	Description string `json:"description"`
	Weight int `json:"weight"`
	Price int `json:"price"`
	Discount int `json:"discount"`
}
