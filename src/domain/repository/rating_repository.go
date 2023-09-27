package repositories

import (
	"context"
	"github.com/Azamjon99/logistics-support-service/src/domain/models"
)

type RatingRepository interface {
	SaveRating(ctx context.Context, rating *models.Rating) error
	UpdateRating(ctx context.Context, rating *models.Rating) error
	GetRating(ctx context.Context, ratingID string) (*models.Rating, error)
	GetRatingByOrder(ctx context.Context, orderID string) (*models.Rating, error)
	}