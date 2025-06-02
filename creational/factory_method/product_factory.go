package factorymethod

// Product defines the interface for all product types in an e-commerce system
type Product interface {
	GetDetails() string

	CalculateShippingCost() float64
}

// ProductFactory defines the interface for creating products
type ProductFactory interface {
	CreateProduct(name string, price float64, extra string) (Product, error)
}

// Inventory manages products in the e-commerce system
type Inventory struct {
	factory ProductFactory
}

func NewInventory(factory ProductFactory) *Inventory {
	return &Inventory{factory: factory}
}

func (i *Inventory) AddProduct(name string, price float64, extra string) (Product, error) {
	return i.factory.CreateProduct(name, price, extra)
}
