package main

import (
	"github.com/labstack/echo"
)

func routes(e *echo.Echo) {
	e.GET("/api/v1/products/:product_id", getProduct)
	e.GET("/api/v1/products", getProducts)
	e.GET("/api/v1/orders/:order_id", getOrder)
	e.GET("/api/v1/orders", getOrders)
	e.POST("/api/v1/order/:product_id", createOrder)
}
