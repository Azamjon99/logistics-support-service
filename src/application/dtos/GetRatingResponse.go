package dtos

import "logistics-support-service/src/domain/models"

type GetRatingResponse struct {
	Rating *models.Rating `json:"rating"`
}

func NewGetRatingResponse(rating *models.Rating) *GetRatingResponse {
	return &GetRatingResponse{
		Rating: rating,
	}
}