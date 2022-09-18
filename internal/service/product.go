package service

import (
	"integration-stripe-go/internal/datastruct"
	"integration-stripe-go/internal/dto"
	"integration-stripe-go/internal/repository"
)

type ProductService interface {
	CreateProduct(product *dto.Product) (*string, error)
}

type productService struct {
	dao repository.DAO
}

func NewProductService(dao repository.DAO) ProductService {
	return &productService{dao: dao}
}

func (p *productService) CreateProduct(product *dto.Product) (*string, error) {
	productInfo := datastruct.Product{
		ID:          product.ID,
		Name:        product.Name,
		Active:      product.Active,
		Description: product.Description,
	}
	productID, err := p.dao.NewProductQuery().CreateProduct(productInfo)
	if err != nil {
		return nil, err
	}
	return productID, nil
}
