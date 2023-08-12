package main

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func getProducts(c echo.Context) error {
	products := retrieveProducts()
	return c.JSON(http.StatusOK, products)
}

func getProduct(c echo.Context) error {
	productId, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid parameter, product id must be an integer"})
	}
	product := retrieveProduct(productId)
	if product != nil {
		return c.JSON(http.StatusOK, product)
	} else {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "product not found"})
	}
}

func getOrders(c echo.Context) error {
	orders := retrieveOrders()
	return c.JSON(http.StatusOK, orders)
}

func getOrder(c echo.Context) error {
	orderId, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid parameter, order id must be an integer"})
	}
	order := retrieveOrder(orderId)
	if order != nil {
		return c.JSON(http.StatusOK, order)
	} else {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "order not found"})
	}
}

func createOrder(c echo.Context) error {
	productId, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid parameter, product id must be an integer"})
	}
	product := retrieveProduct(productId)
	if product != nil {
		insertOrder(productId)
		return c.JSON(http.StatusOK, map[string]string{
			"message": "your order was placed"})
	} else {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "product not found"})
	}
}
