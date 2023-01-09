package random

import (
	"math/rand"

	"github.com/segmentio/ksuid"
)

const maxRandomInt int64 = 999_999

// getKSUIDString returns a K-Sortable Unique IDentifier represented as a string, or an empty string if any error occurred
func getKSUIDString() string {
	id, err := ksuid.NewRandom()
	if err != nil {
		return ""
	}

	return id.String()
}

// getRandomNumber returns a positive pseudo-random number in the interval [1, maxRandomInt]
func getRandomNumber() int64 {
	// No seed is used to make output reproducible
	return rand.Int63n(maxRandomInt) + 1
}
