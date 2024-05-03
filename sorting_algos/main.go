package main

import (
	"fmt"
	"math/rand"
	sorter "sorting/sorting"
	"time"
)

// generate a random int slice of size size and with maximum value for
// every element of max
func generateRandomSlice(size, max int) []int {
	rand.Seed(time.Now().UnixNano())

	result := make([]int, size)

	for i := 0; i < size; i++ {
		result[i] = rand.Intn(max)
	}

	return result
}

var size int = 100000
var slice []int = generateRandomSlice(size, 10000)

func main() {

	// testing the insertion sort algo
	slice1 := make([]int, len(slice))
	copy(slice1, slice)

	time1 := time.Now().UnixMilli()
	sorter.InsertionSort(&slice1)
	time1 = time.Now().UnixMilli() - time1

	// testing the selection sort algo
	slice2 := make([]int, len(slice))
	copy(slice2, slice)

	time2 := time.Now().UnixMilli()
	sorter.IterativeSelectionSort(&slice2)
	time2 = time.Now().UnixMilli() - time2

	// testing the merge sort algo
	slice3 := make([]int, len(slice))
	copy(slice3, slice)

	time3 := time.Now().UnixMilli()
	calls := sorter.MergeSort(&slice3)
	time3 = time.Now().UnixMilli() - time3

	time4 := time.Now().UnixMilli()
	_, err := sorter.Min(0, slice)
	if err != nil {
		panic(err.Error())
	}
	time4 = time.Now().UnixMilli() - time4

	slice4 := make([]int, len(slice))
	copy(slice4, slice)
	time5 := time.Now().UnixMilli()
	sorter.RecursiveSelectionSort(&slice4)
	time5 = time.Now().UnixMilli() - time5

	fmt.Println("iterative selection sort time for size", size, " slice was ", time2, "ms")
	fmt.Println("iterative insertion sort time for size", size, " slice was ", time1, "ms")
	fmt.Println("recursive selection sort time for size", size, " slice was ", time5, "ms")
	fmt.Println("recursive merge sort time for size ", size, " slice was ", time3, "ms")
	fmt.Println("number of recursive calls for merge sort ", calls)
	fmt.Println("recursive min sort time for size", size, " slice was ", time4, "ms")

}
