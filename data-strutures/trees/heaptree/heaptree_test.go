package heaptree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildAnBasicHeapTree(t *testing.T)  {
  t.Run("promotes right child node to root", func(t *testing.T) {
    input := []int{1, 2, 3}

    heap := Build(input)

    expectedHeap := []int{3, 2, 1}

    assert.Equal(t, expectedHeap, heap)
  })


  t.Run("promotes left child node to root", func(t *testing.T) {
    input := []int{2, 3, 1}

    heap := Build(input)

    expectedHeap := []int{3, 2, 1}

    assert.Equal(t, expectedHeap, heap)
  })


  t.Run("Do not change the positions of nodes", func(t *testing.T) {
    input := []int{3, 2, 1}

    heap := Build(input)

    expectedHeap := []int{3, 2, 1}

    assert.Equal(t, expectedHeap, heap)
  })
}

func TestBuildHeapTreeWithTwoLevels(t *testing.T)  {

  t.Run("promotes left leaf node from left child node to root", func(t *testing.T) {
    input := []int{3, 2, 1, 5}

    heap := Build(input)
    // init: [3 2 1 50]
    // 1 step: [3 50 1 2]
    // 2 step: [50 3 1 2]

    expectedHeap := []int{5, 3, 1, 2}

    assert.Equal(t, expectedHeap, heap)
  })

  t.Run("promotes right leaf node from left child node to root", func(t *testing.T) {
    input := []int{5, 2, 1, 3, 6}

    heap := Build(input)
    // init: [5 2 1 3 6]
    // 1 step: [5 6 1 3 2]
    // 2 step: [6 5 1 3 2]

    expectedHeap := []int{6, 5, 1, 3, 2}

    assert.Equal(t, expectedHeap, heap)
  })

  t.Run("promotes right leaf node from right child node to root", func(t *testing.T) {
    input := []int{4, 6, 3, 1, 5, 7}

    heap := Build(input)
    // init: [4 6 3 1 5 2 7]
    // 1 step: [4 6 7 1 5 3]
    // 2 step: [7 6 4 1 5 3]

    expectedHeap := []int{7, 6, 4, 1, 5, 3}

    assert.Equal(t, expectedHeap, heap)
  })

  t.Run("promotes left leaf node from right child node to root", func(t *testing.T) {
    input := []int{4, 6, 3, 1, 5, 2, 7}

    heap := Build(input)
    // init: [4 6 3 1 5 2 7]
    // 1 step: [4 6 7 1 5 2 3]
    // 2 step: [7 6 4 1 5 2 3]

    expectedHeap := []int{7, 6, 4, 1, 5, 2, 3}

    assert.Equal(t, expectedHeap, heap)
  })

  t.Run("promotes left leaf node from right child node to root and past root to leaf node", func(t *testing.T) {
    input := []int{2, 6, 3, 1, 5, 4, 7}

    heap := Build(input)
    // init: [2 6 3 1 5 4 7]
    // 1 step: [2 6 7 1 5 4 3]
    // 2 step: [7 6 2 1 5 4 3]
    // 2 step: [7 6 4 1 5 2 3]

    expectedHeap := []int{7, 6, 4, 1, 5, 2, 3}

    assert.Equal(t, expectedHeap, heap)
  })
}
