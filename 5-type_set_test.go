package belajargogenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Age int

type Number interface {
	~int | int8 | int16 | int32 | int64 | float32 | float64
}

func Min[T Number](first, second T) T {
	if first < second {
		return first
	}

	return second
}

func TestTypeSet(t *testing.T) {
	assert.Equal(t, int(100), Min[int](100, 200))
	assert.Equal(t, int64(100), Min[int64](int64(100), int64(200)))

	// agar bisa digunakan untuk type declaration Age walaupun tipe data sama int, maka gunakan Type Approximation
	// gunakan tanda ~ (tilde) pada tipe data int (atau yg sesuai dengan tipe data dari type declaration yg dibuat)
	assert.Equal(t, Age(100), Min[Age](Age(100), Age(200)))
}

func TestTypeInference(t *testing.T) {
	// dengan type inference tanpa perlu menyebutkan tipedata di dalam kurung, secara otomatis golang tau dari tipe data di parameter
	// di beberapa kasus jika terjadi error tinggal menyebutkan tipe datanya seperti pada TypeSet
	assert.Equal(t, int(100), Min(100, 200))
	assert.Equal(t, int64(100), Min(int64(100), int64(200)))

	// agar bisa digunakan untuk type declaration Age walaupun tipe data sama int, maka gunakan Type Approximation
	// gunakan tanda ~ (tilde) pada tipe data int (atau yg sesuai dengan tipe data dari type declaration yg dibuat)
	assert.Equal(t, Age(100), Min(Age(100), Age(200)))
}
