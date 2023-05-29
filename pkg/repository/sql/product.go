package sql

import (
	"context"
	"fmt"

	"github.com/ReneKroon/ttlcache"
	"github.com/jmoiron/sqlx"

	"api/pkg/entity"
)

type ProductRepository struct {
	db    *sqlx.DB
	cache *ttlcache.Cache
}

func NewProductRepository(db *sqlx.DB, cache *ttlcache.Cache) ProductRepository {
	return ProductRepository{db, cache}
}

func getProductcacheKey(id int) string {
	return fmt.Sprintf("%T:%d", entity.Product{}, id)
}

func (p ProductRepository) GetProduct(ctx context.Context, id int) (entity.Product, error) {
	cachedProduct, exists := p.cache.Get(getProductcacheKey(id))
	if !exists {
		var product entity.Product

		const query = "SELECT id, name, value FROM products where ID = ?"

		err := p.db.Get(&product, query, id)
		if err != nil {
			return entity.Product{}, err
		}

		p.cache.Set(getProductcacheKey(id), product)

		return product, nil
	}

	return cachedProduct.(entity.Product), nil
}

func (p ProductRepository) UpdateProduct(ctx context.Context, id int, product entity.Product) error {
	const query = "UPDATE products SET name = ?, value = ? WHERE id = ? LIMIT 1"
	_, err := p.db.Exec(query, product.Name, product.Value, id)
	if err != nil {
		return err
	}

	p.cache.Remove(getProductcacheKey(id))

	return nil
}

func (p ProductRepository) DeleteProduct(ctx context.Context, id int) error {
	const query = "DELETE FROM products WHERE id = ? LIMIT 1"

	_, err := p.db.Exec(query, id)
	if err != nil {
		return err
	}

	p.cache.Remove(getProductcacheKey(id))

	return nil
}

func (p ProductRepository) CreateProduct(ctx context.Context, product *entity.Product) error {
	const query = "INSERT INTO products (name, value) VALUES(?, ?)"

	res, err := p.db.Exec(query, product.Name, product.Value)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	product.ID = int(id)

	return nil
}
