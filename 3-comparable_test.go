package belajargogenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// comparable tipedata yg bisa dibandingkan menggunakan operator != dan == seperti boolean, number, string, dll
// seperti any, comparable merupakan alias dari tipe data yg bisa untuk perbandingan

func IsSame[T comparable](value1, value2 T) bool {
	if value1 == value2 {
		return value1 == value2
	}

	return false
}

func TestComparable(t *testing.T) {
	assert.True(t, IsSame[string]("Dicki", "Dicki"))
	assert.True(t, IsSame[int](100, 100))
	assert.True(t, IsSame[bool](true, true))
}
