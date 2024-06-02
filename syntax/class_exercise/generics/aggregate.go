package main

import (
	"errors"
	my_basic_go "github.com/hkyz0/my-basic-go"
)

func Sum[T my_basic_go.Number](values ...T) T {
	var sum T
	for _, v := range values {
		sum += v
	}
	return sum
}

func Max[T my_basic_go.RealNumber](values []T) T {
	if len(values) == 0 {
		return values[0]
	}
	max := values[0]
	for _, v := range values {
		if max < v {
			max = v
		}
	}
	return max
}

func Min[T my_basic_go.RealNumber](values ...T) (T, error) {
	if len(values) == 0 {
		var min T
		return min, errors.New("empty values")
	}
	min := values[0]
	for _, v := range values {
		if min < v {
			min = v
		}
	}
	return min, nil
}
