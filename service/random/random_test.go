package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_01_Random_getKSUIDString(t *testing.T) {
	wantLength := 27 // by design of KSUID

	s, err := getKSUIDString()
	require.NoError(t, err)

	assert.Equal(t, len(s), wantLength)
}

func Test_01_Random_getRandomNumber(t *testing.T) {
	n, err := getRandomNumber()
	require.NoError(t, err)

	assert.GreaterOrEqual(t, n, int64(1))
	assert.LessOrEqual(t, n, maxRandomInt)
}
