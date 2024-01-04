package main

import (
	"fmt"
	"os"
	"strconv"

	MergeSort "github.com/DaniloJNS/algorithms-and-data-strutures-in-go/algorithms/sort/mergesort"
)

func main()  {
  args := os.Args[1:] 

  V := make([]int, 0, len(args))

  for _, v := range args {
    i, err := strconv.Atoi(v)
    if err != nil {
      panic(err)
    }
    V = append(V,i)
  }

  fmt.Printf("Input: %+v\n", args)

  MergeSort.Sort(V)

  fmt.Println(V)
}
