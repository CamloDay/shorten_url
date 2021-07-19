package handlers

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_Shorten(t *testing.T) {
	urls := make(map[string]*string, 0)
	original := "http://abc.com"
	short := Shorten(&original, urls, "localhost:8080")
	assert.Equal(t, 58, len(short), "length incorrect, act: %v", short) // short is http://{address}/{uuid}, making it a consistent length
	uuid := strings.TrimPrefix(short, "http://localhost:8080/")
	added := urls[uuid]
	assert.Equal(t, original, *added)
}
