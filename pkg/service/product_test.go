package service_test

import (
	"context"
	"errors"
	"testing"

	"api/pkg/entity"
	"api/pkg/http/api/product"
	"api/pkg/service"
)

// Faker implements ProductRepository for mocking
type faker struct{}

func (faker) CreateProduct(_ context.Context, e *entity.Product) error {
	if e.ID == 1 {
		return nil
	}

	return errors.New("non nil error")
}

func (faker) UpdateProduct(_ context.Context, id int, product entity.Product) error {
	switch id {
	case 1:
		return nil

	default:
		return errors.New("non nil error")
	}

}

func (faker) DeleteProduct(_ context.Context, id int) error {
	switch id {
	case 1:
		return nil

	default:
		return errors.New("non nil error")
	}
}

func (faker) GetProduct(_ context.Context, id int) (entity.Product, error) {
	switch id {
	case 1:
		return entity.Product{ID: 1, Name: "placeholder", Value: 20}, nil

	case 2:
		return entity.Product{}, nil

	default:
		return entity.Product{}, errors.New("non nil error")
	}
}

func TestGetProduct(t *testing.T) {
	productService := service.NewProductService(faker{})

	t.Run("success getting product", func(t *testing.T) {
		product, err := productService.GetProduct(context.Background(), 1)
		if err != nil {
			t.Fatal(err)
		}

		if product.Empty() {
			t.Errorf("GetProduct() = productEmpty: %+v", product)
		}
	})

	t.Run("error from repository", func(t *testing.T) {
		_, err := productService.GetProduct(context.Background(), 3)
		if err == nil {
			t.Error("GetProduct() = should return error")
		}
	})
}

func TestUpdateProduct(t *testing.T) {
	productService := service.NewProductService(faker{})

	t.Run("success updating product", func(t *testing.T) {
		err := productService.UpdateProduct(context.Background(), 1, product.Product{})
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("error from repository", func(t *testing.T) {
		err := productService.UpdateProduct(context.Background(), 3, product.Product{})
		if err == nil {
			t.Error("UpdateProduct() = should return error")
		}
	})
}

func TestDeleteProduct(t *testing.T) {
	productService := service.NewProductService(faker{})

	t.Run("success deleting product", func(t *testing.T) {
		err := productService.DeleteProduct(context.Background(), 1)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("error from repository", func(t *testing.T) {
		err := productService.DeleteProduct(context.Background(), 3)
		if err == nil {
			t.Error("DeleteProduct() = should return error")
		}
	})
}

func TestCreateProduct(t *testing.T) {
	productService := service.NewProductService(faker{})

	t.Run("success creating product", func(t *testing.T) {
		err := productService.CreateProduct(context.Background(), &product.Product{ID: 1})
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("error from repository", func(t *testing.T) {
		err := productService.CreateProduct(context.Background(), &product.Product{})
		if err == nil {
			t.Error("CreateProduct() = should return error")
		}
	})
}
