package random

import (
	"math/rand"

	"github.com/segmentio/ksuid"
)

const maxRandomInt int64 = 999_999

// getKSUIDString returns a K-Sortable Unique IDentifier represented as a string, or an empty string if any error occurred
func getKSUIDString() (string, error) {
	id, err := ksuid.NewRandom()
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

// getRandomNumber returns a positive pseudo-random number in the interval [1, maxRandomInt]
func getRandomNumber() (int64, error) {
	// No seed is used to make output reproducible
	return rand.Int63n(maxRandomInt) + 1, nil
}
