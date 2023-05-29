package api

import (
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"api/config"
	"api/pkg/http/api/product"
	lsql "api/pkg/lib/sql"
	"api/pkg/repository/sql"
	"api/pkg/service"
)

func InjectApiRoutes(e *echo.Echo) {
	g := e.Group("/api")

	g.Use(middleware.BasicAuth(
		func(uname, passwd string, ctx echo.Context) (bool, error) {
			if uname == config.BasicAuthUsername() && passwd == config.BasicAuthPasswd() {
				return true, nil
			}

			return false, nil
		},
	))

	db := lsql.NewDatabase(config.DBPath())

	productHandler := product.NewProductHandler(
		service.NewProductService(
			sql.NewProductRepository(db),
		),
	)

	g.POST("/product/", productHandler.CreateProductHandler)
	g.GET("/product/:id", productHandler.GetProductHandler)
	g.PUT("/product/:id", productHandler.UpdateProductHandler)
	g.DELETE("/product/:id", productHandler.DeleteproductHandler)
}
