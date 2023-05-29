package product_test

import (
	"testing"

	"api/pkg/http/api/product"
)

func TestValidate(t *testing.T) {
	t.Run("valid product", func(t *testing.T) {
		ok := product.Product{ID: 10, Name: "manga", Value: 199.0}.Validate()
		if !ok {
			t.Errorf("product should be valid")
		}
	})

	t.Run("invalid product", func(t *testing.T) {
		ok := product.Product{ID: -1, Value: -69.99}.Validate()
		if ok {
			t.Errorf("product should be invalid")
		}
	})
}
