package main

import (
	"fmt"

	nethttp "github.com/rajenderK7/learn-go/netHttp"
)

const myConst int = 69

var print = fmt.Println

func main() {
	// This is the ultimate Go practice
	// var name string = "Rajender"
	// fmt.Println("This is ", name)

	// var favNumber = 7
	// fmt.Println("The type of this variable is auto inferred -> ", favNumber)
	// fmt.Printf("Type of %v is %T\n", favNumber, favNumber)

	// fmt.Println("This is a constant value: ", myConst)
	// // myConst = 10 // this will give an error

	// firstName := "Katkuri"
	// middleName := ""

	// // len method
	// if len(firstName) > 0 {
	// 	print("First name: ", firstName)
	// }
	// if len(middleName) > 0 {
	// 	print("Middle name: ", middleName)
	// } else {
	// 	print("No middle name")
	// }

	// Array
	// larray.BasicArrayOps()

	// // Slice
	// lslice.BasicSliceOps()
	// lslice.SliceWithInit()
	// lslice.CreateSliceWithMake()

	// // IO
	// lio.PrintNameAndGPA()
	// lio.ReadSliceElements()

	// Map
	// lmap.FindNameOfStudent()

	// Functions
	// if num, isPalindrome := lfunc.Palindrome(); isPalindrome {
	// 	fmt.Printf("%d is a Palindrome\n", num)
	// } else {
	// 	fmt.Printf("%d is NOT a Palindrome\n", num)
	// }

	// ---------------------- BASIC POINTERS
	// var a int = 10
	// names := []string{}
	// print(a, len(names))

	// a := 10
	// b := 11
	// print("a, b =", a, b)
	// var ptr *int = &a
	// ptrB := &b
	// *ptr = 1
	// *ptrB = 2
	// print("a, b =", a, b)
	// -----------------------

	// Pointer
	// a, b := 10, 20
	// print("Before swapping: a, b =", a, b)
	// lpointer.SwapByRef(&a, &b)
	// lpointer.SwapByRefWOTemp(&a, &b)
	// print("After swapping: a, b =", a, b)
	// mp := make(map[string]string)
	// mp["rajender"] = "katkuri"
	// print(mp)
	// lpointer.UnderstandMapModify(mp)
	// print(mp)
	// slc := []int{1, 2, 3, 4, 5}
	// slc := make([]int, 4)
	// print(slc)
	// lpointer.UnderstandSliceModify(slc)
	// print(slc)

	// Struct
	// nums := []int{1, 2, 3, 4, 5}
	// root := lstruct.ConstructBSTfromSlice(nums)
	// root.PrintInorder()
	// fmt.Println()
	// print(root.MorrisInorderTraversal())

	// File IO
	// filename := "batman.txt"
	// I wantedly put the parameter as pointer to string
	// completely unnecessary though
	// if fileCreated := lfileio.CreateCustomFile(&filename); fileCreated {
	// 	print("File created successfully")
	// }
	// lfileio.ReadFromFile(&filename)

	// Net HTTP
	// nethttp.BasicGet("3000")
	// nethttp.BasicPost("3000", "batman")

	// JSON
	nethttp.AddStudents()
	nethttp.JsonEncode()
}
