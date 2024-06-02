package main

import (
	"fmt"
	"strings"
)

func Hello(name string) string {
	if name == "" {
		name = "World"
	} else {
		name = strings.Title(name)
	}
	fmt.Printf("Hello, %s!", name)
	return name
}

func Test(age int) func() {
	fmt.Printf("out: %p", &age)
	println()
	return func() {
		fmt.Printf("before: %p  ", &age)
		age++
		fmt.Printf("after: %p", &age)
		println()
	}
}
