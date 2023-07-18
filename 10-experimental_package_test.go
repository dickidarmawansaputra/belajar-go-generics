package belajargogenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

// gunakan perintah ini
//  go get golang.org/x/exp

func ExperimentalMin[T constraints.Ordered](first, second T) T {
	if first < second {
		return first
	}

	return second
}

func TestExperimentalMin(t *testing.T) {
	assert.Equal(t, 100, ExperimentalMin(100, 200))
	assert.Equal(t, 100.0, ExperimentalMin(100.0, 200.0))
}

func TestExperimentalMaps(t *testing.T) {
	first := map[string]string{
		"Name": "Dicki",
	}
	second := map[string]string{
		"Name": "Dicki",
	}

	assert.True(t, maps.Equal(first, second))
}

func TestExperimentalSlices(t *testing.T) {
	first := []string{"Dicki"}
	second := []string{"Dicki"}

	assert.True(t, slices.Equal(first, second))
}
