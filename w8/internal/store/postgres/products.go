package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"homeworks/w8/internal/models"
	"homeworks/w8/internal/store"
)

func (db *DB) Products() store.ProductsRepository {
	if db.products == nil {
		db.products = NewProductsRepository(db.conn)
	}

	return db.products
}

type ProductsRepository struct {
	conn *sqlx.DB
}

func NewProductsRepository(conn *sqlx.DB) store.ProductsRepository {
	return &ProductsRepository{conn: conn}
}

func (c ProductsRepository) Create(ctx context.Context, product *models.Product) error {
	_, err := c.conn.Exec("INSERT INTO products(name, category_name, brand_name, " +
		"description, weight, price, discount) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		product.Name, product.CategoryName, product.BrandName,
		product.Description, product.Weight, product.Price,
		product.Discount)
	if err != nil {
		return err
	}

	return nil
}

func (c ProductsRepository) All(ctx context.Context) ([]*models.Product, error) {
	products := make([]*models.Product, 0)
	if err := c.conn.Select(&products, "SELECT * FROM products"); err != nil {
		return nil, err
	}

	return products, nil
}

func (c ProductsRepository) ByID(ctx context.Context, id int) (*models.Product, error) {
	product := new(models.Product)
	if err := c.conn.Get(product, "SELECT * FROM products WHERE id=$1", id); err != nil {
		return nil, err
	}

	return product, nil
}

func (c ProductsRepository) Update(ctx context.Context, product *models.Product, id int) error {
	_, err := c.conn.Exec("UPDATE products SET name = $1, " +
		"category_name = $2, brand_name = $3, description = $4," +
		"weight = $5, price = $6, discount = $7 WHERE id = $8",
		product.Name, product.CategoryName, product.BrandName,
		product.Description, product.Weight, product.Price,
		product.Discount, id)
	if err != nil {
		return err
	}

	return nil
}

func (c ProductsRepository) Delete(ctx context.Context, id int) error {
	_, err := c.conn.Exec("DELETE FROM products WHERE id = $1", id)
		if err != nil {
		return err
	}

	return nil
}
