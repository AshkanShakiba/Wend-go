package main

import (
	"github.com/labstack/echo"
)

func routes(e *echo.Echo) {
	e.GET("/api/v1/products/:product_id", getProductHandler)
	e.GET("/api/v1/products", getProductsHandler)
	e.GET("/api/v1/orders/:order_id", getOrderHandler)
	e.GET("/api/v1/orders", getOrdersHandler)
	e.POST("/api/v1/order/:product_id", createOrderHandler)
}
