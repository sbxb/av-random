package rest

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
