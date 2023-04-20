package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBindKey(t *testing.T) {
	expect := "ms_example:f1:a:1:b:2"
	actual := BindKey("f1", map[string]string{"a": "1", "b": "2"})
	assert.Equal(t, expect, actual)
}
