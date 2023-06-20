package lfunc

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var print = fmt.Println

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Palindrome() (int, bool) {
	var num int
	reader := bufio.NewReader(os.Stdin)
	print("Enter a number to check if it is a palindrome:")
	input, err := reader.ReadString('\n')
	CheckNilErr(err)
	parsedInt, err := strconv.ParseInt(strings.TrimSpace(input), 10, 32)
	CheckNilErr(err)
	num = int(parsedInt)
	temp := num
	var rev int = 0
	for num > 0 {
		rev = rev*10 + num%10
		num /= 10
	}
	return temp, rev == temp
}

var GreetUser = func(name string) {
	fmt.Println("Hello", name)
}

var FindMod10 = func(num int) int {
	return num % 10
}
