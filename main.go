package main

import (
	"github.com/labstack/echo"
)

func main() {
	insertProduct("Coffee", 28)
	insertProduct("Coca Cola", 19)

	e := echo.New()

	routes(e)

	e.Logger.Fatal(e.Start(":3000"))
}
