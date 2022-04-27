package go_pkg_demo

import "fmt"

func PrintInfo(value interface{}) interface{} {
	fmt.Println(value)
	return value
}
