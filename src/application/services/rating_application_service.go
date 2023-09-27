package services

import (
	"context"

	"github.com/Azamjon99/logistics-support-service/src/application/dtos"
	pb "github.com/Azamjon99/logistics-support-service/src/application/protos/logistic-support"
	ratingservice "github.com/Azamjon99/logistics-support-service/src/domain/services"
)
 type RatingApplicationService interface {
	CreateRating(ctx context.Context, req *pb.CreateRatingRequest) (*pb.RatingResponse, error)
	UpdateRating(ctx context.Context, req *pb.UpdateRatingRequest) (*pb.RatingResponse, error)
	GetRating(ctx context.Context, req *pb.GetRatingRequest) (*pb.RatingResponse, error)
	GetRatingByOrder(ctx context.Context, req *pb.GetRatingByOrderRequest) (*pb.RatingResponse, error)
 }

 type ratingAppSvcImpl struct {
	ratingSvc ratingservice.RatingService
 }

 func NewRatingApplicationService(ratingSvc ratingservice.RatingService) RatingApplicationService{
	return &ratingAppSvcImpl{
		ratingSvc: ratingSvc,
	}
 }

 func (s *ratingAppSvcImpl) CreateRating(ctx context.Context, req *pb.CreateRatingRequest)(*pb.RatingResponse, error){
	
	result, err := s.ratingSvc.CreateRating(ctx,  req.OrderId, int(req.RatingValue), req.Comment)


	if err != nil {
		return nil, err
	}


    response, err :=dtos.ConvertToRatingResponse(result)
    if err != nil {
        return nil, err
    }
    return response, nil
 }

 func (s *ratingAppSvcImpl) UpdateRating(ctx context.Context, req *pb.UpdateRatingRequest)(*pb.RatingResponse, error){
	
	result, err := s.ratingSvc.UpdateRating(ctx,  req.RatingId, int(req.RatingValue), req.Comment)


	if err != nil {
		return nil, err
	}


    response, err :=dtos.ConvertToRatingResponse(result)
    if err != nil {
        return nil, err
    }
    return response, nil
 }
 

 func (s *ratingAppSvcImpl) GetRating(ctx context.Context, req *pb.GetRatingRequest)(*pb.RatingResponse, error){
	
	result, err := s.ratingSvc.GetRating(ctx,  req.RatingId)


	if err != nil {
		return nil, err
	}


    response, err :=dtos.ConvertToRatingResponse(result)
    if err != nil {
        return nil, err
    }
    return response, nil
 }

 func (s *ratingAppSvcImpl) GetRatingByOrder(ctx context.Context, req *pb.GetRatingByOrderRequest)(*pb.RatingResponse, error){
	
	result, err := s.ratingSvc.GetRatingbyOrder(ctx,  req.OrderId)


	if err != nil {
		return nil, err
	}


    response, err :=dtos.ConvertToRatingResponse(result)
    if err != nil {
        return nil, err
    }
    return response, nil
 }

 


 