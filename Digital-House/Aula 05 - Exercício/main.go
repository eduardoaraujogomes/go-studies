package main

import "fmt"

const (
	SMALL  = "small"
	MEDIUM = "medium"
	LARGE  = "large"
)

type IProduct interface {
	setPrice(price float64)
	getPrice() float64
}

type Product struct {
	price float64
}

func (p *Product) setPrice(price float64) {
	p.price = price
}

func (p *Product) getPrice() float64 {
	return p.price
}

type Small struct {
	Product
}

type Medium struct {
	Product
}

type Large struct {
	Product
}

func newSmallProduct(price float64) IProduct {
	return &Small{
		Product: Product{
			price: price,
		},
	}
}
func newMediumProduct(price float64) IProduct {
	return &Medium{
		Product: Product{
			price: price * 1.03,
		},
	}
}
func newLargeProduct(price float64) IProduct {
	return &Large{
		Product: Product{
			price: (price * 1.06) + 50,
		},
	}
}

func getProductPrice(productType string, value float64) IProduct {
	switch productType {
	case SMALL:
		return newSmallProduct(value)
	case MEDIUM:
		return newMediumProduct(value)
	case LARGE:
		return newLargeProduct(value)
	default:
		return nil
	}
}

func main() {
	fmt.Printf("Total: %.2f \n", getProductPrice(SMALL, 15.23).getPrice())
	fmt.Printf("Total: %.2f \n", getProductPrice(MEDIUM, 21.15).getPrice())
	fmt.Printf("Total: %.2f \n", getProductPrice(LARGE, 33.27).getPrice())
}
