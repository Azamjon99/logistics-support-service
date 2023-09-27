package rating

import (
	"context"
	"github.com/Azamjon99/logistics-support-service/src/domain/models"
	repositories "github.com/Azamjon99/logistics-support-service/src/domain/repository"

	"gorm.io/gorm"
)

const(

  ratingTable = "support.ratings"
)

type ratingrepoImpl struct {
	db *gorm.DB
}

func NewRatingRepository(db *gorm.DB) repositories.RatingRepository{
	return &ratingrepoImpl{
		db:db,
	}
}

func (r *ratingrepoImpl)SaveRating(ctx context.Context, rating *models.Rating) error{
	result := r.db.WithContext(ctx).Table(ratingTable).Create(&rating)

	if result.Error != nil{
		return result.Error
	}

	return nil;
}

func (r *ratingrepoImpl)UpdateRating(ctx context.Context, rating *models.Rating) error{
	result := r.db.WithContext(ctx).Table(ratingTable).Save(&rating)

	if result.Error != nil{
		return result.Error
	}

	return nil;
}



func (r *ratingrepoImpl)GetRating(ctx context.Context, ratingID string)(*models.Rating, error){
	var rating *models.Rating
	result := r.db.WithContext(ctx).Table(ratingTable).First(&rating, "rating_id = ?", ratingID)

	if result.Error != nil {
		return nil, result.Error
	}

	return rating, nil;
}

func (r *ratingrepoImpl)GetRatingByOrder(ctx context.Context, orderID string)(*models.Rating, error){
	var rating *models.Rating
	result := r.db.WithContext(ctx).Table(ratingTable).First(&rating, "order_id = ?", orderID)

	if result.Error != nil {
		return nil, result.Error
	}

	return rating, nil;
}