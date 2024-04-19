package sorting

import (
	"errors"
)

// func merge sort without slicing (TODO slicing)
// it returns the number of recursive calls
func MergeSort(A *[]int) int {
	X := make([]int, len(*A))
	return mergeSort(A, 0, len(*A), X)
}

func mergeSort(A *[]int, i, f int, X []int) int {
	if f-i > 1 {
		m := (i + f) / 2
		n1 := mergeSort(A, i, m, X)
		n2 := mergeSort(A, m, f, X)
		merge(A, i, m, f, X)
		return 1 + n1 + n2
	} else {
		return 0
	}
}

func merge(A *[]int, i, m, f int, X []int) {
	i1 := i
	i2 := m
	k := 0
	for i1 < m && i2 < f {
		if (*A)[i1] <= (*A)[i2] {
			X[k] = (*A)[i1]
			i1++
		} else {
			X[k] = (*A)[i2]
			i2++
		}
		k++
	}
	if i1 < m {
		for j := i1; j < m; j++ {
			X[k] = (*A)[j]
			k++
		}
	} else {
		for j := i2; j < f-1; j++ {
			X[k] = (*A)[j]
			k++
		}
	}
	for k := 0; k < f-i-1; k++ {
		(*A)[i+k] = X[k]
	}
}

func InsertionSort(A *[]int) {
	j := 0
	for k, x := range (*A)[1 : len(*A)-1] {
		j = k - 1
		for j >= 0 && (*A)[j] > x {
			(*A)[j+1] = (*A)[j]
			j = j - 1
		}
		(*A)[j+1] = x
	}
}

func Min(k int, A []int) (int, error) {
	if len(A) == 0 {
		return -1, errors.New("the list A is empty")
	}
	return minRec(A), nil
}

func minRec(A []int) int {
	if len(A) == 1 {
		return (A)[0]
	}

	m := len(A) / 2
	a := minRec(A[:m])
	b := minRec(A[m:])

	if a >= b {
		return b
	} else {
		return a
	}
}

func IterativeSelectionSort(A *[]int) {
	n := len(*A)
	var min, aux int

	for k := 0; k < n-1; k++ {
		min = k
		for j, y := range (*A)[k+1 : n-1] {
			if y < (*A)[min] {
				min = j
			}
		}
		aux = (*A)[min]
		(*A)[min] = (*A)[k]
		(*A)[k] = aux
	}
}

func RecursiveSelectionSort(A *[]int) {
	recSelectionSort(0, A)
}

func recSelectionSort(k int, A *[]int) {
	if len(*A)-k > 1 {
		min := k
		for j, y := range (*A)[k+1 : len(*A)-1] {
			if y < (*A)[min] {
				min = j
			}
		}
		aux := (*A)[min]
		(*A)[min] = (*A)[k]
		(*A)[k] = aux
		recSelectionSort(k+1, A)
	}
}
