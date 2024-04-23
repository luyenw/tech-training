package main

import "fmt"

type Product struct {
	name     string
	price    float64
	discount float64
}

type ProductBuilder struct {
	name     string
	price    float64
	discount float64
}

func NewProductBuilder() *ProductBuilder {
	return &ProductBuilder{}
}

func (pb *ProductBuilder) SetName(name string) *ProductBuilder {
	pb.name = name
	return pb
}

func (pb *ProductBuilder) SetPrice(price float64) *ProductBuilder {
	pb.price = price
	return pb
}

func (pb *ProductBuilder) SetDiscount(discount float64) *ProductBuilder {
	pb.discount = discount
	return pb
}

func (pb *ProductBuilder) Build() *Product {
	return &Product{
		name:     pb.name,
		price:    pb.price,
		discount: pb.discount,
	}
}

func main() {
	builder := NewProductBuilder().
		SetName("Laptop").
		SetPrice(1000).
		SetDiscount(100)
	product := builder.Build()

	fmt.Printf("%+v", *product)
}
