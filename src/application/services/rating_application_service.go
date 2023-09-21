package services

import (
	"context"
	"errors"
	"logistics-support-service/src/application/dtos"
	ratingservice "logistics-support-service/src/domain/services"
)
 type RatingApplicationService interface {
	CreateRating(ctx context.Context, ratingID, orderID string, ratingValue int, comment string) (*dtos.GetRatingResponse, error)
 }

 type ratingAppSvcImpl struct {
	ratingSvc ratingservice.RatingService
 }

 func NewRatingApplicationService(ratingSvc ratingservice.RatingService) RatingApplicationService{
	return &ratingAppSvcImpl{
		ratingSvc: ratingSvc,
	}
 }

 func (s *ratingAppSvcImpl) CreateRating(ctx context.Context, ratingID string, orderID string, ratingValue int, comment string)(*dtos.GetRatingResponse, error){
	if ratingID ==""{
		return nil, errors.New("Rating id is required")
	}
	if orderID ==""{
		return nil, errors.New("orderID is required")
	}
	if comment ==""{
		return nil, errors.New("comment  is required")
	}
	if ratingValue ==0{
		return nil, errors.New("ratingValue is required")
	}
	if ratingID ==""{
		return nil, errors.New("Rating id is required")
	}
	result, err := s.ratingSvc.CreateRating(ctx, ratingID, orderID, ratingValue, comment)


	if err != nil {
		return nil, err
	}

	return dtos.NewGetRatingResponse(result), nil
 }