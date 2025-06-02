package factorymethod

import (
	"errors"
	"strings"
)

var (
	ErrInvalidProductPrice = errors.New("invalid product price")
	ErrInvalidProductName  = errors.New("invalid product name")
)

type baseProduct struct {
	price float64
	name  string
}

func (bf *baseProduct) validatePrice() error {
	if bf.price <= 0 {
		return ErrInvalidProductPrice
	}

	return nil
}

func (bf *baseProduct) validateName() error {
	bf.name = strings.Trim(bf.name, " ")

	if bf.name == "" {
		return ErrInvalidProductName
	}

	return nil
}
