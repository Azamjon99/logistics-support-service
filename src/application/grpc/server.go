package grpc

import (
	"context"

	pb "github.com/Azamjon99/logistics-support-service/src/application/protos/logistic-support"
	service "github.com/Azamjon99/logistics-support-service/src/application/services"
)




type Server struct{
	pb.RatingServiceServer
	ratingApp	service.RatingApplicationService
}

func NewServer(ratingApp	service.RatingApplicationService) *Server {
		return &Server{
			ratingApp: ratingApp,
		 }
}

func (s *Server)  CreateRating(ctx context.Context, r *pb.CreateRatingRequest) (*pb.RatingResponse, error){
	return s.ratingApp.CreateRating(ctx, r)
	
}

func (s *Server)  UpdateRating(ctx context.Context, r *pb.UpdateRatingRequest) (*pb.RatingResponse, error){
	return s.ratingApp.UpdateRating(ctx, r)
	
}

func (s *Server)  GetRating(ctx context.Context, r *pb.GetRatingRequest) (*pb.RatingResponse, error){
	return s.ratingApp.GetRating(ctx, r)
	
}
func (s *Server)  GetRatingByOrder(ctx context.Context, r *pb.GetRatingByOrderRequest) (*pb.RatingResponse, error){
	return s.ratingApp.GetRatingByOrder(ctx, r)
	
}