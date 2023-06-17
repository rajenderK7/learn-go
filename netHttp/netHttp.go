package nethttp

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/rajenderK7/learn-go/lfunc"
)

var print = fmt.Println

func BasicGet(port string) {
	// Here I used a node server🥲 Switch to Go!!
	url := "http://localhost:" + port + "/bro"
	response, err := http.Get(url)
	lfunc.CheckNilErr(err)
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	print(response.Status)
	lfunc.CheckNilErr(err)
	print(string(data))
}

func BasicPost(port string, name string) {
	var url = "http://localhost:" + port + "/greet/" + name
	reqBody := strings.NewReader(`{ "message": "Bro" }`)
	response, err := http.Post(url, "application/json", reqBody)
	lfunc.CheckNilErr(err)
	defer response.Body.Close()
	data, _ := io.ReadAll(response.Body)
	print(string(data))
}
