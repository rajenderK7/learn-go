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
