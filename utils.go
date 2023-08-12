package main

func getProductObj(product Product) map[string]interface{} {
	return map[string]interface{}{
		"id":    product.id,
		"name":  product.name,
		"price": product.price}
}

func getOrderObj(order Order) map[string]interface{} {
	return map[string]interface{}{
		"id":      order.id,
		"product": order.product,
		"date":    order.date.Format("2006-01-02 15:04:05")}
}
