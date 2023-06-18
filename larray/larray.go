package larray

import "fmt"

func BasicArrayOps() {
	var myArray [4]int
	myArray[0] = 6
	myArray[1] = 9
	myArray[2] = 6
	myArray[3] = 9
	fmt.Println("My array is: ", myArray)
	fmt.Println("Length of myArray: ", len(myArray))
}

// Accepting an array without mentioning the size
// is effectively accepting a slice.
func PassingArrayAsParam(arr [4]int) {
	n := len(arr)
	if n > 0 {
		arr[n-1] = -1
	}
	fmt.Println("From the function Pass by value: ", arr)
}

func PassingArrayAsPointer(arr *[4]int) {
	n := len(arr)
	if n > 1 {
		arr[n-1] = -1
	}
	fmt.Println("From the function Pointer: ", *arr)

}
