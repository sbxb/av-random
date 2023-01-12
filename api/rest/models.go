package rest

import "github.com/sbxb/av-random/models"

type GenerateRequest struct {
	Type   string `json:"type"`
	Length int    `json:"length"`
}

type GenerateResponse struct {
	ID string `json:"generation_id"`
}

// type RetrieveRequest struct {
// }

type RetrieveResponse struct {
	ID           string `json:"generation_id"`
	RandomNumber string `json:"random_number"`
}

func convRandomEntityToRetrieveResponse(re models.RandomEntity) RetrieveResponse {
	return RetrieveResponse{
		ID:           re.GenerationID,
		RandomNumber: re.RandomValue,
	}
}

func (gr GenerateRequest) validate() bool {
	isTypeValid := false
	for _, t := range models.RandomValueTypes {
		if gr.Type == t {
			isTypeValid = true
			break
		}
	}

	return isTypeValid && gr.Length > 0
}
