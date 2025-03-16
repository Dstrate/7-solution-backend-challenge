package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/Dstrate/7-solution-backend-challenge-3/beef/proto"
	env "github.com/joho/godotenv"
	"google.golang.org/grpc"

	"github.com/gofiber/fiber/v2"
)

type server struct {
	pb.UnimplementedBeefServiceServer
}

func startGrpcServer() {
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		log.Fatalf("GRPC_PORT is not set.")
	}
	listener, _ := net.Listen("tcp", ":"+grpcPort)
	grpcServer := grpc.NewServer()
	pb.RegisterBeefServiceServer(grpcServer, &server{})
	fmt.Printf("Start GRPC Server Port %s\n", grpcPort)
	grpcServer.Serve(listener)
}

func startFiberServer() {
	apiPort := os.Getenv("GRPC_PORT")
	if apiPort == "" {
		log.Fatalf("GRPC_PORT is not set.")
	}
	app := fiber.New()
	// register route
	app.Get("/beef/summary", BeefSummaryService)
	fmt.Printf("Start API Server Port %s\n", apiPort)
	app.Listen(":" + apiPort)
}

func main() {
	err := env.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// สร้าง 2 server แบบ api กับ grpc
	go startGrpcServer()
	go startFiberServer()

	<-ctx.Done()
	fmt.Println("Shutdown Server...")
}
