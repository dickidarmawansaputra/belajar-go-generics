package belajargogenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee interface {
	GetName() string
}

func GetName[T Employee](parameter T) string {
	return parameter.GetName()
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

type VicePresident interface {
	GetName() string
	GetVicePresidentName() string
}

type MyVicePresident struct {
	Name string
}

func (m *MyVicePresident) GetName() string {
	return m.Name
}

func (m *MyVicePresident) GetVicePresidentName() string {
	return m.Name
}

func TestGetName(t *testing.T) {
	// mirip inheritance di OOP, cuma karna di golang bukan berbasis OOP maka tidak disebutkan secara langsung
	assert.Equal(t, "Dicki", GetName[Manager](&MyManager{Name: "Dicki"}))
	assert.Equal(t, "Dicki", GetName[VicePresident](&MyVicePresident{Name: "Dicki"}))
}
