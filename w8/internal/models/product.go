package models

type Product struct {
	Id int `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	CategoryName string `json:"category_name" db:"category_name"`
	BrandName string `json:"brand_name" db:"brand_name"`
	Description string `json:"description" db:"description"`
	Weight int `json:"weight" db:"weight"`
	Price int `json:"price" db:"price"`
	Discount int `json:"discount" db:"discount"`
}
