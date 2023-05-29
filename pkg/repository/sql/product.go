package sql

import (
	"context"

	"github.com/jmoiron/sqlx"

	"api/pkg/entity"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) ProductRepository {
	return ProductRepository{db}
}

func (p ProductRepository) GetProduct(ctx context.Context, id int) (entity.Product, error) {
	var product entity.Product

	const query = "SELECT id, name, value FROM products where ID = ?"

	err := p.db.Get(&product, query, id)
	return product, err
}

func (p ProductRepository) UpdateProduct(ctx context.Context, id int, product entity.Product) error {
	// TODO : update db

	return nil
}

func (p ProductRepository) DeleteProduct(ctx context.Context, id int) error {
	// TODO : delete db

	return nil
}

func (p ProductRepository) CreateProduct(ctx context.Context, product entity.Product) error {
	// TODO : create db

	return nil
}
