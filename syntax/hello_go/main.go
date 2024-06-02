package main

import "fmt"

func main() {
	s1 := []int{2, 4, 6, 8, 10}
	s2 := s1[:3]
	fmt.Printf("s2: %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))

	s3 := s1[2:3]
	fmt.Printf("s3: %v, len: %d, cap: %d\n111", s3, len(s3), cap(s3))
}
