package calculation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		a := 1
		b := 2
		expected := 3
		actual := Add(a, b)
		assert.Equal(t, expected, actual, "results should be equal")
	})
}
