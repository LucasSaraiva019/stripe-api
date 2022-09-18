package repository

import (
	"fmt"
	"integration-stripe-go/internal/datastruct"

	"github.com/stripe/stripe-go/v73"
)

type ProductQuery interface {
	CreateProduct(product datastruct.Product) (*string, error)
}

type productQuery struct{}

func (p *productQuery) CreateProduct(product datastruct.Product) (*string, error) {
	params := stripe.ProductParams{
		Name:        &product.Name,
		Description: &product.Description,
		Active:      &product.Active,
	}
	stripeProduct, err := stripeQb().Products.New(&params)
	if err != nil {
		return nil, fmt.Errorf("cannot create a new product: %v", err)
	}

	return &stripeProduct.ID, nil
}
