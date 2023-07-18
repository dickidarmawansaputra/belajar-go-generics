package belajargogenerics

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, value := range bag {
		fmt.Println(value)
	}
}

func TestBagString(t *testing.T) {
	names := Bag[string]{"Dicki", "Darmawan", "Saputra"}
	PrintBag[string](names)
}

func TestBagInt(t *testing.T) {
	number := Bag[int]{1, 2, 3}
	PrintBag[int](number)
}
