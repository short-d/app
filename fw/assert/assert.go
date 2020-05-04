package assert

import (
	"testing"

	testify "github.com/stretchr/testify/assert"
)

func Equal(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	return testify.Equal(t, expected, actual, msgAndArgs)
}

func NotEqual(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	return testify.NotEqual(t, expected, actual, msgAndArgs)
}

func SameElements(t *testing.T, expected, actual interface{}) bool {
	return testify.ElementsMatch(t, expected, actual)
}
