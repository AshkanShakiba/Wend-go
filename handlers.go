package main

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func getProductsHandler(c echo.Context) error {
	products := getProductsService()
	return c.JSON(http.StatusOK, products)
}

func getProductHandler(c echo.Context) error {
	productId, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid parameter, product id must be an integer"})
	}
	product := getProductService(productId)
	if product != nil {
		return c.JSON(http.StatusOK, product)
	} else {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "product not found"})
	}
}

func getOrdersHandler(c echo.Context) error {
	orders := getOrdersService()
	return c.JSON(http.StatusOK, orders)
}

func getOrderHandler(c echo.Context) error {
	orderId, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid parameter, order id must be an integer"})
	}
	order := getOrderService(orderId)
	if order != nil {
		return c.JSON(http.StatusOK, order)
	} else {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "order not found"})
	}
}

func createOrderHandler(c echo.Context) error {
	productId, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid parameter, product id must be an integer"})
	}
	product := getProductService(productId)
	if product != nil {
		createOrderService(productId)
		return c.JSON(http.StatusOK, map[string]string{
			"message": "your order was placed"})
	} else {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "product not found"})
	}
}
