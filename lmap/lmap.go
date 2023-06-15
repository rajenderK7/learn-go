package lmap

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var print = fmt.Println

// This func ouputs the name of the student
// given their Roll No which is unique to each
// student. It internally uses a map where
// Roll No is int and Name is string
func FindNameOfStudent() {
	students := make(map[int]string)
	reader := bufio.NewReader(os.Stdin)
	var prompt string = "Enter \"Roll_No Name\", and \"x \"x to exit"
	print(prompt)
	var roll int
	var name string
	for {
		input, _ := reader.ReadString('\n')
		fields := strings.Fields(input)
		if len(fields) > 2 {
			print("Please enter only Roll_No and Name")
			continue
		}
		if fields[0] == "x" && fields[1] == "x" {
			break
		}
		roll_conv, _ := strconv.ParseInt(strings.TrimSpace(fields[0]), 10, 32)
		roll = int(roll_conv)
		name = fields[1]
		students[roll] = name
	}
	for key, val := range students {
		print(key, val)
	}
	delete(students, 11) // given a key 11
	print(students)
}
