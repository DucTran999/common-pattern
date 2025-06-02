package factorymethod_test

import (
	"log"
	factorymethod "patterns/creational/factory_method"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PhysicalFactory(t *testing.T) {
	type testCase struct {
		testName     string
		name         string
		price        float64
		weight       string
		productStr   string
		shippingCost float64
		expectedErr  error
	}

	testcases := []testCase{
		{
			testName:    "invalid name",
			name:        "",
			expectedErr: factorymethod.ErrInvalidProductName,
		},
		{
			testName:    "invalid price",
			name:        "product A",
			price:       0,
			expectedErr: factorymethod.ErrInvalidProductPrice,
		},
		{
			testName:    "invalid weight",
			name:        "product A",
			price:       2,
			weight:      "a",
			expectedErr: factorymethod.ErrInvalidProductWeight,
		},
		{
			testName:     "valid product",
			name:         "product A",
			price:        2,
			weight:       "2.6",
			productStr:   "Physical Product: product A, Price: $2.00, Weight: 2.60 kg",
			shippingCost: 13,
			expectedErr:  nil,
		},
	}

	pf := factorymethod.NewInventory(&factorymethod.PhysicalProductFactory{})

	for _, tc := range testcases {
		t.Run(tc.testName, func(t *testing.T) {
			product, err := pf.AddProduct(tc.name, tc.price, tc.weight)

			if err != nil {
				assert.Equal(t, tc.expectedErr, err)
			} else {
				assert.Equal(t, tc.productStr, product.GetDetails())
				log.Println(product.CalculateShippingCost())
				assert.InEpsilon(t, tc.shippingCost, product.CalculateShippingCost(), 1)
			}
		})
	}
}

func Test_DigitalFactory(t *testing.T) {
	type testCase struct {
		testName     string
		name         string
		price        float64
		downloadLink string
		productStr   string
		shippingCost float64
		expectedErr  error
	}

	testcases := []testCase{
		{
			testName:    "invalid name",
			name:        "",
			expectedErr: factorymethod.ErrInvalidProductName,
		},
		{
			testName:    "invalid price",
			name:        "product A",
			price:       0,
			expectedErr: factorymethod.ErrInvalidProductPrice,
		},
		{
			testName:     "invalid download link",
			name:         "product A",
			price:        2,
			downloadLink: "",
			expectedErr:  factorymethod.ErrInvalidProductLink,
		},
		{
			testName:     "valid product",
			name:         "product A",
			price:        2,
			downloadLink: "url",
			productStr:   "Digital Product: product A, Price: $2.00, Download: url",
			shippingCost: 13,
			expectedErr:  nil,
		},
	}

	pf := factorymethod.NewInventory(&factorymethod.DigitalProductFactory{})

	for _, tc := range testcases {
		t.Run(tc.testName, func(t *testing.T) {
			product, err := pf.AddProduct(tc.name, tc.price, tc.downloadLink)

			if err != nil {
				assert.Equal(t, tc.expectedErr, err)
			} else {
				assert.Equal(t, tc.productStr, product.GetDetails())
				log.Println(product.CalculateShippingCost())
				assert.InEpsilon(t, tc.shippingCost, product.CalculateShippingCost(), 1)
			}
		})
	}
}
