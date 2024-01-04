package mergesort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortAnSimpleArray(t *testing.T)  {
  t.Run("must swap the first element with the last", func(t *testing.T) {
    arr := []int{3, 2, 1}

    Sort(arr)

    expectedArray := []int{1, 2, 3}

    assert.Equal(t, expectedArray, arr)
  })

  t.Run("rerange array", func(t *testing.T) {
    arr := []int{3, 2, 1, 5, 10}

    Sort(arr)

    expectedArray := []int{1, 2, 3, 5, 10}

    assert.Equal(t, expectedArray, arr)
  })

  t.Run("Rerange array with duplicates elements", func(t *testing.T) {
    arr := []int{10, 2, 1, 5, 10}

    Sort(arr)

    expectedArray := []int{1, 2, 5, 10, 10}

    assert.Equal(t, expectedArray, arr)
  })
}
