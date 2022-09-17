package repository

import (
	"fmt"
	"integration-stripe-go/internal/datastruct"
	"strconv"

	"github.com/stripe/stripe-go/v73"
)

type ProductQuery interface {
	CreateProduct(product datastruct.Product) (*int64, error)
}

type productQuery struct{}

func (p *productQuery) CreateProduct(product datastruct.Product) (*int64, error) {
	params := stripe.ProductParams{
		Name:        &product.Name,
		Description: &product.Description,
		Active:      &product.Active,
	}
	stripeProduct, err := stripeQb().Products.New(&params)
	if err != nil {
		return nil, fmt.Errorf("cannot create a new product: %v", err)
	}

	id, err := strconv.Atoi(stripeProduct.ID)
	if err != nil {
		return nil, fmt.Errorf("cannot convert the product id: %v", err)
	}

	return stripe.Int64(int64(id)), nil
}
