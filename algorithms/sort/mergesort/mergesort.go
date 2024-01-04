package mergesort

func partition(input *[]int, aux *[]int, leftCursor int, rigthCursor int) {
	if leftCursor < rigthCursor {
		middle := leftCursor + (rigthCursor - leftCursor)/2

		partition(input, aux, leftCursor, middle)
		partition(input, aux, middle + 1, rigthCursor)
		merge(input, aux, leftCursor, middle, rigthCursor)
	}
}


func merge(input *[]int, aux *[]int, leftCursor int, middle int, rigthCursor int) {
	i1 := leftCursor
	i2 := middle + 1
	k := leftCursor

  // Insere elementos dos subarray
  // A -> (Leftcusror .. middle)
  // B -> (middle +1 .. rigthCursor)
  // no vetor auxliar de forma ordenada,
  // enquanto ambos os subvetores contirem elementos.
  // Ex:
  // Input -> [3, 5, 2], leftCursor: 0, rigthCursor: 2, logo
  // A -> [3, 5]
  // B -> [2]
  // Sendo Aux -> []
  // Apos a primeira e unica iteracao temos
  // A -> [3, 5], B -> [] e Aux -> [2]
  // Os elementos nao mergeados de A serao apendados no vetor Aux. Assim,
  // Aux -> [2, 3, 5]
	for i1 <= middle && i2 <= rigthCursor {
		if (*input)[i1] >= (*input)[i2] {
			(*aux)[k] = (*input)[i2]
			i2++
		} else {
			(*aux)[k] = (*input)[i1]
			i1++
		}
		k++
	}

  // Copia para o vetor auxiliar os elementos restante do subarray que contem elementos
	if i1 > middle {
	  copy((*aux)[k:], (*input)[i2:rigthCursor+1])
	} else {
	  copy((*aux)[k:], (*input)[i1:middle+1])
	}

  copy((*input)[leftCursor:rigthCursor+1], (*aux)[leftCursor:rigthCursor+1])
}

func Sort(input []int) {
	aux := make([]int, len(input))
	partition(&input, &aux, 0, len(input) - 1)
}
