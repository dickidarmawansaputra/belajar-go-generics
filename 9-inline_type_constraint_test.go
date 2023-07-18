package belajargogenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func FindMin[T interface{ int | int64 | float64 }](first, second T) T {
	if first < second {
		return first
	}

	return second
}

func TestFindMin(t *testing.T) {
	assert.Equal(t, 100, FindMin[int](100, 100))
	assert.Equal(t, int64(100), FindMin(int64(100), int64(100)))
	assert.Equal(t, 100.0, FindMin(100.0, 100.0))
}

func GetFirst[T []E, E any](slice T) E {
	first := slice[0]
	return first
}

func TestGetFirst(t *testing.T) {
	names := []string{
		"Dicki", "Darmawan", "Saputra",
	}
	getFirst := GetFirst[[]string, string](names)
	assert.Equal(t, "Dicki", getFirst)
}
