package product_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"

	mock_product "api/mocks"
	"api/pkg/http/api/product"
	"api/pkg/lib/mock"
)

var requestJson = `{"name": "bicycle", "value": 123.45}`

func TestProductHandler_CreateProductHandler(t *testing.T) {
	requestJson := `{"name": "bicycle", "value": 123.45}`

	ctrl := gomock.NewController(t)
	mockService := mock_product.NewMockProductService(ctrl)

	handler := product.NewProductHandler(mockService)

	t.Run("success creating product", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api", strings.NewReader(requestJson))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		echoContext := echo.New().NewContext(req, rec)

		mockService.EXPECT().CreateProduct(gomock.Any(), mock.MatchedBy(func(matcher interface{}) bool {
			product := matcher.(*product.Product)
			return product.Name == "bicycle" && product.Value == 123.45
		})).Return(nil)

		err := handler.CreateProductHandler(echoContext)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("error from repository creating product", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api", strings.NewReader(requestJson))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		echoContext := echo.New().NewContext(req, rec)

		mockService.EXPECT().CreateProduct(gomock.Any(), mock.MatchedBy(func(matcher interface{}) bool {
			product := matcher.(*product.Product)
			return product.Name == "bicycle" && product.Value == 123.45
		})).Return(errors.New("non nil error"))

		err := handler.CreateProductHandler(echoContext)
		if err == nil {
			t.Fatal(err)
		}
	})
}

func TestProductHandler_GetProductHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := mock_product.NewMockProductService(ctrl)

	handler := product.NewProductHandler(mockService)

	t.Run("success getting product", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := echo.New().NewContext(req, rec)
		c.SetPath("/products/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		mockService.EXPECT().GetProduct(gomock.Any(), 1).
			Return(product.Product{ID: 1, Name: "bicycle", Value: 123.45}, nil)

		err := handler.GetProductHandler(c)
		if err != nil {
			t.Fatal(err)
		}

		if http.StatusOK != rec.Code {
			t.Errorf("status not ok: %d", rec.Code)
		}
	})

	t.Run("error from repository getting product", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		echoContext := echo.New().NewContext(req, rec)

		mockService.EXPECT().GetProduct(gomock.Any(), 1).
			Return(product.Product{}, errors.New("non nil error"))

		err := handler.GetProductHandler(echoContext)
		if err != nil {
			t.Fatal(err)
		}

		if http.StatusNotFound != rec.Code {
			t.Errorf("status not ok: %d", rec.Code)
		}
	})
}

func TestProductHandler_UpdateProductHandler(t *testing.T) {
	requestJson := `{"name": "bicycle", "value": 123.45}`

	ctrl := gomock.NewController(t)
	mockService := mock_product.NewMockProductService(ctrl)

	handler := product.NewProductHandler(mockService)

	t.Run("success updating product", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/api", strings.NewReader(requestJson))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := echo.New().NewContext(req, rec)
		c.SetPath("/products/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		mockService.EXPECT().UpdateProduct(gomock.Any(), 1, product.Product{Name: "bicycle", Value: 123.45}).
			Return(nil)

		err := handler.UpdateProductHandler(c)
		if err != nil {
			t.Fatal(err)
		}

		if http.StatusOK != rec.Code {
			t.Errorf("status not ok: %d", rec.Code)
		}
	})

	t.Run("error from repository updating product", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/api", strings.NewReader(requestJson))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := echo.New().NewContext(req, rec)
		c.SetPath("/products/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		mockService.EXPECT().UpdateProduct(gomock.Any(), 1, product.Product{Name: "bicycle", Value: 123.45}).
			Return(errors.New("non nil error"))

		err := handler.UpdateProductHandler(c)
		if err == nil {
			t.Fatal(err)
		}
	})
}

func TestProductHandler_DeleteProductHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := mock_product.NewMockProductService(ctrl)

	handler := product.NewProductHandler(mockService)

	t.Run("success creating product", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/api", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := echo.New().NewContext(req, rec)
		c.SetPath("/products/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		mockService.EXPECT().DeleteProduct(gomock.Any(), 1).
			Return(nil)

		err := handler.DeleteproductHandler(c)
		if err != nil {
			t.Fatal(err)
		}

		if http.StatusOK != rec.Code {
			t.Errorf("status not ok: %d", rec.Code)
		}
	})

	t.Run("error from repository creating product", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/api", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := echo.New().NewContext(req, rec)
		c.SetPath("/products/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		mockService.EXPECT().DeleteProduct(gomock.Any(), 1).
			Return(errors.New("non nil errror"))

		err := handler.DeleteproductHandler(c)
		if err == nil {
			t.Fatal(err)
		}
	})
}
