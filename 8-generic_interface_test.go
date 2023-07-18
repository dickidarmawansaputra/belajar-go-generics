package belajargogenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetterSeeter[T any] interface {
	GetValue() T
	SetValue(value T)
}

func ChangeValue[T any](param GetterSeeter[T], value T) T {
	param.SetValue(value)
	return param.GetValue()
}

type MyData[T any] struct {
	Value T
}

func (d *MyData[T]) GetValue() T {
	return d.Value
}

func (d *MyData[T]) SetValue(value T) {
	d.Value = value
}

func TestGenericInterface(t *testing.T) {
	myData := MyData[string]{}

	result := ChangeValue[string](&myData, "Dicki")

	assert.Equal(t, "Dicki", result)
	assert.Equal(t, "Dicki", myData.Value)
}
