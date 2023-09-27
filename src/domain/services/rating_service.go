package services

import (
	"context"
	"github.com/Azamjon99/logistics-support-service/src/domain/models"
	repositories "github.com/Azamjon99/logistics-support-service/src/domain/repository"
	"time"
)

type RatingService interface {
	CreateRating(ctx context.Context,  orderID string, ratingValue int, comment string) (*models.Rating, error)
	UpdateRating(ctx context.Context, ratingID string, ratingValue int, comment string) (*models.Rating, error)
	GetRating(ctx context.Context, ratingID string) (*models.Rating, error)
	GetRatingbyOrder(ctx context.Context, orderID string) (*models.Rating, error)
	}

type ratingSvcImpl struct {
	ratingRepo repositories.RatingRepository
}

func NewRatingSvc( ratingRepo repositories.RatingRepository) RatingService{
	return &ratingSvcImpl{
		ratingRepo: ratingRepo,
	}
}

func (s *ratingSvcImpl) CreateRating(ctx context.Context,  orderID string, ratingValue int, comment string)(*models.Rating, error){
	rating :=  &models.Rating{
		OrderID: orderID,
		Rating: ratingValue,
		Comment: comment,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	err:= s.ratingRepo.SaveRating(ctx, rating)

	if err != nil {
		return nil, err
	}

	return rating, nil;
}

func (s *ratingSvcImpl) UpdateRating(ctx context.Context, ratingID string, ratingValue int, comment string) (*models.Rating, error) {
	rating := &models.Rating{
		ID:       ratingID,
		Rating:   ratingValue,
		Comment:  comment,
		UpdatedAt: time.Now().UTC(),
	}

	err := s.ratingRepo.UpdateRating(ctx, rating)

	if err != nil {
		return nil, err
	}

	return rating, nil
}

func (s *ratingSvcImpl) GetRating(ctx context.Context, ratingID string)(*models.Rating, error){


	rating,	err:= s.ratingRepo.GetRating(ctx, ratingID)

	if err != nil {
		return nil, err
	}

	return rating, nil;
}

func (s *ratingSvcImpl) GetRatingbyOrder(ctx context.Context, OrderID string)(*models.Rating, error){


	rating,	err:= s.ratingRepo.GetRatingByOrder(ctx, OrderID)

	if err != nil {
		return nil, err
	}

	return rating, nil;
}