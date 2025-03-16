package main

import (
	"context"

	pb "github.com/Dstrate/7-solution-backend-challenge-3/beef/proto"
	"github.com/gofiber/fiber/v2"

	"github.com/Dstrate/7-solution-backend-challenge-3/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// For GRPC Server
func (s *server) BeefSummaryService(ctx context.Context, req *pb.BeefSummaryRequest) (*pb.BeefSummaryResponse, error) {

	beef := services.GetNewBeefService()
	countBeefs, err := beef.GetBeefSummary()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error: %v", err.Error())
	}
	return &pb.BeefSummaryResponse{
		Beef: countBeefs,
	}, nil
}

// For API Server
func BeefSummaryService(c *fiber.Ctx) error {

	beef := services.GetNewBeefService()
	countBeefs, err := beef.GetBeefSummary()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{
		"beef": countBeefs,
	})
}
