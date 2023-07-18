package belajargogenerics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// biasanya nama dalam kurung cuma 1 karakter biar tidak membingungkan
// jika tipe data interface kosong, semua tipe data bisa
// tipe data any = alias dari interface kosong
// scopenya cuma bisa di func yg didefinisikan type paramaternya
// dengan gunakan generic bisa membuat func yg bisa dipake dengan tipe data berbeda, yg dalamnya dengan logic yg sama
// jadi cukup 1 kode dengan fitur ini bisa dipake dengan tipe data berbeda
// sebutan tipe data = constraints contoh tipe data any alias interface kosong
func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestLength(t *testing.T) {
	var resultString = Length[string]("Dicki")
	assert.Equal(t, "Dicki", resultString)

	var resultInt = Length[int](100)
	assert.Equal(t, 100, resultInt)
}
