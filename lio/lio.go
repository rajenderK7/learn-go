package lio

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var print = fmt.Println

// Read and Write from and to Std IO
func PrintNameAndGPA() {
	reader := bufio.NewReader(os.Stdin)
	print("Enter your name:")
	name, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	// trim space is required because the reader is using
	// newline char as a delim so the delim is also read by the buffer
	// and we do not wish to have it in our value
	name = strings.TrimSpace(name)
	print("Enter your GPA:")
	gpaInput, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	// trim space is again important here
	gpa, _ := strconv.ParseFloat(strings.TrimSpace(gpaInput), 32)
	fmt.Printf("Name: %s\tGPA: %.2f\n", name, gpa)
}

func ReadSliceElements() {
	reader := bufio.NewReader(os.Stdin)
	print("Enter elements of the slice:")
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	elements := strings.Fields(input)
	// create a slice
	nums := make([]int, 0)
	for _, ele := range elements {
		val, _ := strconv.ParseInt(ele, 10, 32)
		nums = append(nums, int(val))
	}
	print("Number of elements: ", len(nums))
	print(nums)
}
