package main

import (
	"context"
	"log"
	"net/http"

	pb "github.com/campoy/links/microservices-grpc/repository/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	var config struct {
		Address    string `default:"localhost:8081"`
		Repository string `default:"localhost:8080"`
	}
	if err := envconfig.Process("GATEWAY", &config); err != nil {
		log.Fatal(err)
	}

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterRepositoryHandlerFromEndpoint(context.Background(), mux, config.Repository, opts)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("serving gRPC endpoint from %s on REST %s", config.Repository, config.Address)
	log.Fatal(http.ListenAndServe(config.Address, mux))
}
