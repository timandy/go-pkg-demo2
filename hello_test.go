package go_pkg_demo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrintInfo(t *testing.T) {
	info := PrintInfo(33)
	assert.Equal(t, 33, info)
}
