package heapsort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortAnSimpleArray(t *testing.T)  {
  t.Run("must swap the first element with the last", func(t *testing.T) {
    input := []int{3, 2, 1}

    sortedArray := Sort(input)

    expectedArray := []int{1, 2, 3}

    assert.Equal(t, expectedArray, sortedArray)
  })

  t.Run("rerange array", func(t *testing.T) {
    input := []int{3, 2, 1, 5, 10}

    sortedArray := Sort(input)

    expectedArray := []int{1, 2, 3, 5, 10}

    assert.Equal(t, expectedArray, sortedArray)
  })

  t.Run("rerange array with duplicates elements", func(t *testing.T) {
    input := []int{10, 2, 1, 5, 10}

    sortedArray := Sort(input)

    expectedArray := []int{1, 2, 5, 10, 10}

    assert.Equal(t, expectedArray, sortedArray)
  })
}
