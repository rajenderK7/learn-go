package nethttp

import (
	"encoding/json"

	"github.com/rajenderK7/learn-go/lfunc"
)

type Student struct {
	Name                 string   `json:"name"`
	Roll                 string   `json:"roll"`
	Department           string   `json:"dept"`
	ProgrammingLanguages []string `json:"languages,omitempty"`
	FavNumber            int      `json:"favNumber"`
	FavPerson            string   `json:"-"`
}

func CreateStudent(name string, roll string, dept string, languages []string, favNumber int, favPerson string) *Student {
	return &Student{name, roll, dept, languages, favNumber, favPerson}
}

var students []Student

func AddStudents() {
	// Some dummy data dude!!
	for i := 0; i < 3; i++ {
		// CreateStudent returns the address of a new student
		// but, students slice is holding actual objects hence we need to
		// derefernce the address present in the pointer returned by CreateStudent
		students = append(students, *CreateStudent("Batman", "7", "Knight", []string{"C++", "Go", "Kotlin"}, 69, "Cat Woman"))
	}
}

func JsonEncode() {
	// json, err := json.Marshal(students) // Marshal will give raw json (similar to JSON.stringify)
	json, err := json.MarshalIndent(students, "", "\t") // Second argument `prefix` prefixes the provided string for each line of the indented JSON response
	lfunc.CheckNilErr(err)
	print(string(json))
}
