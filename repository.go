package main

import (
	"time"
)

var products []Product
var orders []Order

func insertProduct(name string, price int) {
	product := Product{id: len(orders), name: name, price: price}
	addProduct(product)
}

func addProduct(product Product) {
	products = append(products, product)
}

func retrieveProduct(productId int) map[string]interface{} {
	for i := 0; i < len(products); i++ {
		if products[i].id == productId {
			return getProductObj(products[i])
		}
	}
	return nil
}

func retrieveProducts() []map[string]interface{} {
	var products_list []map[string]interface{}
	for i := 0; i < len(products); i++ {
		products_list = append(products_list, getProductObj(products[i]))
	}
	return products_list
}

func insertOrder(productId int) {
	order := Order{id: len(orders), product: productId, date: time.Now()}
	addOrder(order)
}

func addOrder(order Order) {
	orders = append(orders, order)
}

func retrieveOrder(orderId int) map[string]interface{} {
	for i := 0; i < len(orders); i++ {
		if orders[i].id == orderId {
			return getOrderObj(orders[i])
		}
	}
	return nil
}

func retrieveOrders() []map[string]interface{} {
	var orders_list []map[string]interface{}
	for i := 0; i < len(orders); i++ {
		orders_list = append(orders_list, getOrderObj(orders[i]))
	}
	return orders_list
}
