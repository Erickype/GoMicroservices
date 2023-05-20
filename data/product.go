package data

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io"
	"regexp"
	"time"
)

// Product defines the structure of a product
type Product struct {
	ID          int     `json:"id,omitempty"`
	Name        string  `json:"name,omitempty" validate:"required"`
	Description string  `json:"description,omitempty"`
	Price       float32 `json:"price,omitempty" validate:"gt=0"`
	SKU         string  `json:"sku,omitempty" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// FromJSON decodes a json object to Product
func (p *Product) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

func (p *Product) Validate() error {
	validate := validator.New()
	err := validate.RegisterValidation("sku", validateSKU)
	if err != nil {
		return err
	}

	err = validate.Struct(p)
	if err != nil {
		return err
	}
	return nil
}

func validateSKU(fl validator.FieldLevel) bool {
	// sku format: asd-asd-asd
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}

	return true
}

// Products is a collection of Product
type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func GetProducts() Products {
	return productsList
}

func AddProduct(product *Product) {
	product.ID = getNextID()
	productsList = append(productsList, product)
}

func UpdateProduct(id int, product *Product) error {
	actualProduct, position, err := findProduct(id)
	if err != nil {
		return err
	}
	product.ID = actualProduct.ID
	productsList[position] = product
	return nil
}

func DeleteProduct(id int) error {
	_, position, err := findProduct(id)
	if err != nil {
		return err
	}
	productsList = append(productsList[:position], productsList[position+1:]...)
	return nil
}

var ErrProductNotFound = fmt.Errorf("product not found")

func findProduct(id int) (*Product, int, error) {
	for i, product := range productsList {
		if product.ID == id {
			return product, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

func getNextID() int {
	lastIndex := productsList[len(productsList)-1].ID
	return lastIndex + 1
}

var productsList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc232",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "dld493",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
