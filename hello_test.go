package go_pkg_demo

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintInfo(t *testing.T) {
	fmt.Println("==========", runtime.GOOS, runtime.GOARCH)
	info := PrintInfo(33)
	assert.Equal(t, 33, info)
}
