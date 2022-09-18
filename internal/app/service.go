package app

import (
	"integration-stripe-go/internal/service"
	pb "integration-stripe-go/pkg"
)

type MicroserviceServer struct {
	pb.UnimplementedStripeServer
	productService service.ProductService
}

func NewMicroservice(
	productService service.ProductService,
) *MicroserviceServer {
	return &MicroserviceServer{
		productService: productService,
	}
}
