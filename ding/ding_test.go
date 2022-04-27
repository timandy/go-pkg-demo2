package ding

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDing(t *testing.T) {
	info := Ding(33)
	assert.Equal(t, 33, info)
}
