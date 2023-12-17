package heaptree

func left_node(parent_idx int) int {
  return (2 * parent_idx) + 1
}

func right_node(parent_idx int) int {
  return (2 * parent_idx) + 2
}

func Swap(V []int, a int, b int)  {
  V[a], V[b] = V[b], V[a]
}

func Heapfy(V []int, size int, current_idx int) {
  P, L, R := current_idx, left_node(current_idx), right_node(current_idx)

  P = current_idx

  if L < size && V[L] > V[P]  {
    P = L
  }

  if R < size && V[R] > V[P]  {
    P = R
  }

  if P != current_idx {
    Swap(V, P, current_idx)
    Heapfy(V, size, P)
  }
}

func Build(V []int) []int {
  for i := (len(V) - 1) / 2; i >= 0; i-- {
    Heapfy(V, len(V), i)
  }
  return V
}
