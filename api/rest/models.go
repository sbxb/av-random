package rest

import "github.com/sbxb/av-random/models"

// type GenerateRequest struct {
// }

type GenerateResponse struct {
	ID string `json:"generation_id"`
}

// type RetrieveRequest struct {
// }

type RetrieveResponse struct {
	ID           string `json:"generation_id"`
	RandomNumber int64  `json:"random_number"`
}

func convRandomEntityToRetrieveResponse(re models.RandomEntity) RetrieveResponse {
	return RetrieveResponse{
		ID:           re.GenerationID,
		RandomNumber: re.RandomValue,
	}
}
