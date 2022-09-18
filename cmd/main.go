package main

import (
	"context"
	"integration-stripe-go/internal/app"
	"integration-stripe-go/internal/repository"
	"integration-stripe-go/internal/service"
	pb "integration-stripe-go/pkg"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	//STRIPE
	client, err := repository.NewStripe()
	if err != nil {
		return
	}

	// Preparing config file
	viper.AddConfigPath("../config")
	viper.SetConfigName("config")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalln("cannot read from a config")
	}
	// Interceptors
	//grpcOpts := app.GrpcInterceptor()
	//httpOpts := app.HttpInterceptor()

	// Register all services
	dao := repository.NewDAO(client)
	productService := service.NewProductService(dao)
	// Starting gRPC server
	go func() {
		listener, err := net.Listen("tcp", "localhost:8081")
		if err != nil {
			log.Fatalln(err)
		}

		grpcServer := grpc.NewServer()
		pb.RegisterStripeServer(grpcServer, app.NewMicroservice(
			productService))
		log.Printf("server listening at %v", listener.Addr())
		err = grpcServer.Serve(listener)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	// Stating HTTP server
	mux := runtime.NewServeMux()
	err = pb.RegisterStripeHandlerServer(context.Background(), mux, app.NewMicroservice(
		productService,
	))
	if err != nil {
		log.Println("cannot register this service")
	}
	log.Printf("server HTTP listening at 127.0.0.1:8080")
	log.Fatalln(http.ListenAndServe(":8080", mux))
}
