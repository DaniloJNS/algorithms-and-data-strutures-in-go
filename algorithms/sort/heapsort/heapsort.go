package heapsort

import (
	"fmt"

	HeapTree "github.com/DaniloJNS/algorithms-and-data-strutures-in-go/data-strutures/trees/heaptree"
)


func swap(arr []int, a int, b int)  {
  arr[a], arr[b] = arr[b], arr[a]
}

func Sort(values []int) []int {
  heap := HeapTree.Build(values)

  fmt.Printf("Heap: %+v\n", heap)

  last_idx := len(heap) - 1

  for i := 0; i < len(heap) - 1; i++ {
    swap(heap, 0, last_idx - i)
    HeapTree.Heapfy(heap, last_idx - i, 0)
  }

  fmt.Printf("Sorted: %+v\n", heap)

  return heap
}
