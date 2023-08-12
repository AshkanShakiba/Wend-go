package main

import (
	"time"
)

type Product struct {
	id    int
	name  string
	price int
}

type Order struct {
	id      int
	product int
	date    time.Time
}
