package main

func getProductsService() []map[string]interface{} {
	products := retrieveProducts()
	return products
}

func getProductService(productId int) map[string]interface{} {
	product := retrieveProduct(productId)
	return product
}

func createProductService(name string, price int) {
	insertProduct(name, price)
}

func getOrdersService() []map[string]interface{} {
	orders := retrieveOrders()
	return orders
}

func getOrderService(orderId int) map[string]interface{} {
	order := retrieveOrder(orderId)
	return order
}

func createOrderService(productId int) {
	insertOrder(productId)
}
