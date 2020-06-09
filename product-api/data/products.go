package data

import (
	"fmt"
)

// ErrProductNotFound is an error raised when a product can not be found in the database
var ErrProductNotFound = fmt.Errorf("Product not found")

// Product defines the structure for an API product.
// Product defines the structure for an API product
// swagger:model
type Product struct {
	// the id for the product
	//
	// required: false
	// min: 1
	ID int `json:"id"` // Unique identifier for the product

	// the name for this poduct
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// the description for this poduct
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`

	// the price for the product
	//
	// required: true
	// min: 0.01
	Price float32 `json:"price" validate:"required,gt=0"`

	// the SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU string `json:"sku" validate:"sku"`
}

type Products []*Product

type ProductError struct {
	errorMsg string
}

// GetProduct returns the Product structure of the given id. If the id is not in the "database"
// then we return an error.
func GetProduct(id int) (*Product, error) {

	product, _, err := FindProduct(id)

	return product, err
}

func FindProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

func UpdateProduct(prodToUpdate *Product) error {
	_, pos, err := FindProduct(prodToUpdate.ID)

	if err != nil {
		return err
	}

	productList[pos] = prodToUpdate

	return nil
}

func AddProduct(p *Product) {
	maxID := productList[len(productList)-1].ID
	p.ID = maxID + 1
	productList = append(productList, p)
}

func DeleteProduct(id int) error {
	_, i, err := FindProduct(id)

	if err != nil {
		return err
	}

	// That's linear complexity. Oof.
	productList = append(productList[:i], productList[i+1])

	return nil
}

func GetProducts() Products {
	return productList
}

var productList = Products{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "bcd-bcd-bcd",
	},
	&Product{
		ID:          2,
		Name:        "Esspresso",
		Description: "Short strong coffee without milk",
		Price:       1.99,
		SKU:         "abc-abc-abc",
	},
}
