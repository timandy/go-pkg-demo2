package ding

import "fmt"

func Ding[T any](value T) T {
	fmt.Println("ding", value)
	return value
}
