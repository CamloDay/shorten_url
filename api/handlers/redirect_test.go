package handlers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MapOriginal(t *testing.T) {
	urls := make(map[string]*string, 0)
	original := "abc"
	urls["def"] = &original

	act, err := MapOriginal(urls, "def")
	assert.NoError(t, err)
	assert.Equal(t, "abc", *act)
}

func Test_MapOriginalNotFound(t *testing.T) {
	urls := make(map[string]*string, 0)

	_, err := MapOriginal(urls, "def")
	assert.EqualError(t, err, StatusDescriptions[StatusNotFound])
}
