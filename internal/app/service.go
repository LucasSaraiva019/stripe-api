package app

import (
	"integration-stripe-go/internal/service"
	pb "integration-stripe-go/pkg"
)

type MicroserviceServer struct {
	productService service.ProductService
	pb.UnimplementedStripeServer
}

func NewMicroservice(
	productService service.ProductService,
) *MicroserviceServer {
	return &MicroserviceServer{
		productService: productService,
	}
}
