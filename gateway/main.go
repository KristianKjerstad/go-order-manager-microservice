package main

import (
	common "commons"
	"log"
	"net/http"

	pb "commons/api"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var httpAddr = common.EnvString("HTTP_ADDR", ":8070")
var orderServiceAddr = "localhost:2000"

func main() {

	conn, err := grpc.Dial(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	log.Println("dialing orders service at", orderServiceAddr)
	c := pb.NewOrderServiceClient(conn)

	mux := http.NewServeMux()
	handler := NewHandler(c)

	handler.registerRoutes(mux)
	log.Printf("Starting HTTP server at %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start http server", err)

	}

}
