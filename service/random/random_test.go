package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_01_Random_getKSUIDString(t *testing.T) {
	wantLength := 27 // by design of KSUID

	s := getKSUIDString()
	t.Log(s)

	assert.Equal(t, len(s), wantLength)
}

func Test_01_Random_getRandomNumber(t *testing.T) {
	n := getRandomNumber()
	t.Log(n)

	assert.GreaterOrEqual(t, n, int64(1))
	assert.LessOrEqual(t, n, maxRandomInt)
}
