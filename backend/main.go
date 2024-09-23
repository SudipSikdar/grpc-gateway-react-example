package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/SudipSikdar/go-grpc-gateway-react/pb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	pb.UnimplementedGrpcServiceServer
}

func main() {

	go startGRPCServer()
	startRestGateway()
}

func (s *server) GetMessage(ctx context.Context, in *pb.GetMessageRequest) (*pb.GetMessageResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.GetMessageResponse{Message: "Hello " + in.GetName()}, nil
}

func startGRPCServer() {

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGrpcServiceServer(s, &server{})

	log.Printf("Starting gRPC server listening on %s", listener.Addr().String())

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func startRestGateway() {
	log.Printf("Starting HTTP gateway listening on :8081")
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := pb.RegisterGrpcServiceHandlerFromEndpoint(context.Background(), mux, ":8080", opts)

	if err != nil {
		log.Fatalf("failed to start HTTP gateway: %v", err)
	}

	// Enable CORS
	corsHanlder := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(mux)

	log.Printf("Starting HTTP gateway listening on :8081")
	err = http.ListenAndServe(":8081", corsHanlder)
	if err != nil {
		log.Fatalf("failed to start HTTP gateway: %v", err)
	}

}
