package rest

import (
	"encoding/json"
	"fmt"
	"io"
)

func parsePostGenerateBody(body io.ReadCloser) (GenerateRequest, error) {
	req := GenerateRequest{}

	dec := json.NewDecoder(body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(&req); err != nil {
		return req, fmt.Errorf("Can not decode GenerateRequest with error: %w", err)
	}

	if !req.validate() {
		return req, fmt.Errorf("Can not validate GenerateRequest: %v", req)
	}

	return req, nil
}
