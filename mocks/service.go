// Code generated by MockGen. DO NOT EDIT.
// Source: handler.go

// Package mock_product is a generated GoMock package.
package mock_product

import (
	product "api/pkg/http/api/product"
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	reflect "reflect"
)

// MockProductService is a mock of ProductService interface
type MockProductService struct {
	ctrl     *gomock.Controller
	recorder *MockProductServiceMockRecorder
}

// MockProductServiceMockRecorder is the mock recorder for MockProductService
type MockProductServiceMockRecorder struct {
	mock *MockProductService
}

// NewMockProductService creates a new mock instance
func NewMockProductService(ctrl *gomock.Controller) *MockProductService {
	mock := &MockProductService{ctrl: ctrl}
	mock.recorder = &MockProductServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProductService) EXPECT() *MockProductServiceMockRecorder {
	return m.recorder
}

// CreateProduct mocks base method
func (m *MockProductService) CreateProduct(arg0 context.Context, arg1 *product.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProduct indicates an expected call of CreateProduct
func (mr *MockProductServiceMockRecorder) CreateProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockProductService)(nil).CreateProduct), arg0, arg1)
}

// UpdateProduct mocks base method
func (m *MockProductService) UpdateProduct(ctx context.Context, id int, product product.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", ctx, id, product)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProduct indicates an expected call of UpdateProduct
func (mr *MockProductServiceMockRecorder) UpdateProduct(ctx, id, product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockProductService)(nil).UpdateProduct), ctx, id, product)
}

// DeleteProduct mocks base method
func (m *MockProductService) DeleteProduct(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProduct indicates an expected call of DeleteProduct
func (mr *MockProductServiceMockRecorder) DeleteProduct(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockProductService)(nil).DeleteProduct), ctx, id)
}

// GetProduct mocks base method
func (m *MockProductService) GetProduct(ctx context.Context, id int) (product.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProduct", ctx, id)
	ret0, _ := ret[0].(product.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduct indicates an expected call of GetProduct
func (mr *MockProductServiceMockRecorder) GetProduct(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockProductService)(nil).GetProduct), ctx, id)
}
