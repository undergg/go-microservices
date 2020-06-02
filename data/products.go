package data

import (
	"encoding/json"
	"io"
	"time"
)

// Product defines the structure for an API product.
// We use struct tags to change the variable name in the JSON objects.
// We can omit some fields that we don't want to return as part of the
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

type ProductError struct {
	errorMsg string
}

// Implement the Error() function to agree with the error interface.
func (e *ProductError) Error() string {
	return e.errorMsg
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

	return nil, -1, &ProductError{"Product not found"}
}

func UpdateProduct(id int, prodToUpdate *Product) error {
	_, pos, err := FindProduct(id)

	if err != nil {
		return err
	}

	productList[pos] = prodToUpdate
	// Make sure the id of the product passed in the request body is the same :).
	// TODO: Maybe do some validation for it?
	productList[pos].ID = id

	return nil
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

var productList = Products{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Esspresso",
		Description: "Short strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
