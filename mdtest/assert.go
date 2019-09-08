package mdtest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Equal(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	return assert.Equal(t, expected, actual, msgAndArgs)
}

func NotEqual(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	return assert.NotEqual(t, expected, actual, msgAndArgs)
}

func SameElements(t *testing.T, expected, actual interface{}) bool {
	return assert.ElementsMatch(t, expected, actual)
}
