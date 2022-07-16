package service

import (
	"fmt"
	"github.com/golang-rest-shop-backend/pkg/database"
	. "github.com/golang-rest-shop-backend/pkg/structs"
)

type ExchangeRateAPIResponse struct {
	Success   bool   `json:"success"`
	Timestamp int    `json:"timestamp"`
	Base      string `json:"base"`
	Date      string `json:"date"`
	Rates     struct {
		BGN float64 `json:"BGN"`
		CAD float64 `json:"CAD"`
		CHF float64 `json:"CHF"`
		EUR float64 `json:"EUR"`
		GBP float64 `json:"GBP"`
		USD float64 `json:"USD"`
	} `json:"rates"`
}

func GetAllProducts(currency string) ([]Product, error) {
	products, err := database.GetAllProducts()
	if err != nil {
		return nil, fmt.Errorf("failed to get all products with error: %s\n", err)
	}

	for i := range products {

		if &products[i] != nil {
			return nil, err
		}
	}

	return products, nil
}

func GetProductById(id string, currency string) (*Product, error) {
	product, err := database.GetProductById(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find such product error: %s\n", err)
	}

	return product, nil
}

func AddProduct(product *Product) (string, error) {
	productId, err := database.AddProduct(product)
	if err != nil {
		return "", err
	}

	return productId, nil
}

func UpdateProduct(product *Product) error {
	err := database.UpdateProduct(product)
	if err != nil {
		return err
	}

	return nil
}

func DeleteProduct(productId string) error {
	if err := database.DeleteProduct(productId); err != nil {
		return err
	}

	return nil
}
