package dtos

import (
	"encoding/json"
	pb "github.com/Azamjon99/logistics-support-service/src/application/protos/logistic-support"

)

func ConvertToRatingResponse(result interface{}) (*pb.RatingResponse, error) {
    resultJSON, err := json.Marshal(result)
    if err != nil {
        return nil, err
    }

    var resultMap map[string]interface{}
    if err := json.Unmarshal(resultJSON, &resultMap); err != nil {
        return nil, err
    }


    rating := &pb.Rating{
    }

    response := &pb.RatingResponse{
        Rating: rating,
    }

    return response, nil
}