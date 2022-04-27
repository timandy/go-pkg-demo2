package go_pkg_demo

import "fmt"

func PrintInfo[T any](value T) T {
	fmt.Println(value)
	return value
}
