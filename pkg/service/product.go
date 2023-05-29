package service

import (
	"context"

	"api/pkg/entity"
	productModel "api/pkg/http/api/product"
)

type ProductRepository interface {
	CreateProduct(context.Context, *entity.Product) error
	UpdateProduct(ctx context.Context, id int, product entity.Product) error
	DeleteProduct(ctx context.Context, id int) error
	GetProduct(ctx context.Context, id int) (entity.Product, error)
}

type ProductService struct {
	productRepository ProductRepository
}

func NewProductService(pr ProductRepository) ProductService {
	return ProductService{pr}
}

func (p ProductService) GetProduct(ctx context.Context, id int) (productModel.Product, error) {
	product, err := p.productRepository.GetProduct(ctx, id)
	if err != nil {
		return productModel.Product{}, err
	}

	return productModel.NewProduct(product), nil
}

func (p ProductService) UpdateProduct(ctx context.Context, id int, product productModel.Product) error {
	return p.productRepository.UpdateProduct(ctx, id, productModel.GetProduct(product))
}

func (p ProductService) DeleteProduct(ctx context.Context, id int) error {
	return p.productRepository.DeleteProduct(ctx, id)
}

func (p ProductService) CreateProduct(ctx context.Context, product *productModel.Product) error {
	entityProduct := productModel.GetProduct(*product)

	err := p.productRepository.CreateProduct(ctx, &entityProduct)
	if err != nil {
		return err
	}

	product.ID = entityProduct.ID

	return nil
}
