package factorymethod

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	ErrInvalidProductWeight = errors.New("invalid weight for physical product")
)

// PhysicalProductFactory creates physical products
type PhysicalProductFactory struct{}

func (f *PhysicalProductFactory) CreateProduct(name string, price float64, extra string) (Product, error) {
	baseProduct := baseProduct{
		name:  name,
		price: price,
	}

	if err := baseProduct.validateName(); err != nil {
		return nil, err
	}

	if err := baseProduct.validatePrice(); err != nil {
		return nil, err
	}

	weight, err := strconv.ParseFloat(extra, 64)
	if err != nil || weight <= 0 {
		return nil, ErrInvalidProductWeight
	}

	return &PhysicalProduct{name: name, price: price, weight: weight}, nil
}

// PhysicalProduct represents a tangible item like a t-shirt or laptop
type PhysicalProduct struct {
	name   string
	price  float64
	weight float64 // in kilograms
}

func (p *PhysicalProduct) GetDetails() string {
	return fmt.Sprintf("Physical Product: %s, Price: $%.2f, Weight: %.2f kg", p.name, p.price, p.weight)
}

func (p *PhysicalProduct) CalculateShippingCost() float64 {
	// Simple shipping cost: $5 per kg
	return p.weight * 5.0
}
