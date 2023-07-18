package belajargogenerics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data[T any] struct {
	First  T
	Second T
}

// jika tidak butuh generic type gunakan tanda _
func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return d.First
}

func TestData(t *testing.T) {
	data := Data[string]{
		First:  "Dicki",
		Second: "DS",
	}

	fmt.Println(data)
}

func TestGenericMethod(t *testing.T) {
	data := Data[string]{
		First:  "Dicki",
		Second: "DS",
	}

	assert.Equal(t, "Dicki", data.ChangeFirst("Dicki"))
	assert.Equal(t, "Hello Dicki", data.SayHello("Dicki"))
}
