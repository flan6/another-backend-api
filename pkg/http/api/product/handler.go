//go:generate go run -mod=mod github.com/golang/mock/mockgen -source=$GOFILE -aux_files=product=model.go -destination=../../../../mocks/service.go
package product

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/context"
)

type ProductService interface {
	CreateProduct(context.Context, *Product) error
	UpdateProduct(ctx context.Context, id int, product Product) error
	DeleteProduct(ctx context.Context, id int) error
	GetProduct(ctx context.Context, id int) (Product, error)
}

type ProductHandler struct {
	productService ProductService
}

func NewProductHandler(ps ProductService) ProductHandler {
	return ProductHandler{ps}
}

func (p ProductHandler) CreateProductHandler(c echo.Context) error {
	var productRequest CreateProductRequest

	err := json.NewDecoder(c.Request().Body).Decode(&productRequest)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "could not create product, invalid payload")
	}

	product := Product{
		Name:  productRequest.Name,
		Value: productRequest.Value,
	}

	ok := product.Validate()
	if !ok {
		return c.String(http.StatusBadRequest, "invalid product payload")
	}

	err = p.productService.CreateProduct(context.Background(), &product)
	if err != nil {
		c.Logger().Error(err)
		return errors.New("could not create product")
	}

	return c.String(http.StatusOK, fmt.Sprintf("CreateProductHandler - id: %d", product.ID))
}

func (p ProductHandler) GetProductHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusNotFound, "could not find product")
	}

	product, err := p.productService.GetProduct(context.Background(), id)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusNotFound, "could not find product")
	}

	return c.JSON(http.StatusOK, product)
}

func (p ProductHandler) UpdateProductHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusNotFound, "could not find product")
	}

	var requestProduct UpdateProductRequest

	body := c.Request().Body
	if body == nil {
		return errors.New("empty body")
	}

	err = json.NewDecoder(body).Decode(&requestProduct)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "invalid product payload")
	}

	product := Product{
		Name:  requestProduct.Name,
		Value: requestProduct.Value,
	}

	err = p.productService.UpdateProduct(context.Background(), id, product)
	if err != nil {
		c.Logger().Error(err)
		return errors.New("could not update product")
	}

	return c.String(http.StatusOK, "UpdateProductHandler")
}

func (p ProductHandler) DeleteproductHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusNotFound, "could not find product")
	}

	err = p.productService.DeleteProduct(context.Background(), id)
	if err != nil {
		c.Logger().Error(err)
		return errors.New("unnable to delete product")
	}

	return c.String(http.StatusOK, "DeleteProductHandler")
}
