package lslice

import "fmt"

var print = fmt.Println

func BasicSliceOps() {
	// Create a slice of integers
	// var nums [6]int // this is an array
	var nums []int // slice
	// the above decleration is same as var nums = []int{}
	nums = append(nums, 10, 11, 12, 13)
	print(nums)
	print(len(nums))
}

func SliceWithInit() {
	var names = []string{"rajender", "kothi", "pandi"}
	print("Slice with Initialization")
	print(len(names))
	names = append(names, "bruh")
	print(len(names))
}

func CreateSliceWithMake() {
	// size := 3
	names := make([]string, 0)
	names = append(names, "apple")
	print("Length of names: ", len(names))
	print("names: ", names)
}
