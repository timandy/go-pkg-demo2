package ding

import "fmt"

func Ding(value interface{}) interface{} {
	fmt.Println("ding", value)
	return value
}
