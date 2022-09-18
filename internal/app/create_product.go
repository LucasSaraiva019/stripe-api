package app

import (
	"context"
	"integration-stripe-go/internal/dto"
	pb "integration-stripe-go/pkg"
)

func (m *MicroserviceServer) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	product := dto.Product{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Active:      req.GetActive(),
	}
	productID, err := m.productService.CreateProduct(&product)
	if err != nil {
		return nil, err
	}
	return &pb.CreateProductResponse{Id: *productID}, nil
}
