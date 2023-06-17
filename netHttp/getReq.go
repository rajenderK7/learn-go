package nethttp

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/rajenderK7/learn-go/lfunc"
)

type Address struct {
	Street string `json:"street"`
	Suite  string `json:"suite"`
	City   string `json:"city"`
	Zip    string `json:"zipcode"`
	Geo    Geo    `json:"geo"`
}

type Geo struct {
	Latitude  string `json:"lat"`
	Longitude string `json:"lng"`
}

type Company struct {
	Name    string `json:"name"`
	TagLine string `json:"catchPhrase"`
	Bs      string `json:"bs"`
}

type User struct {
	Id        int     `json:"id"`
	Name      string  `json:"string"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	Address   Address `json:"address"`
	Phone     string  `json:"phone"`
	Portfolio string  `json:"website"`
	Company   Company `json:"company"`
}

const jsonPlaceholderEndpoint string = "https://jsonplaceholder.typicode.com/users"

var users = []User{
	{
		1,
		"Leanne Graham",
		"Bret",
		"Sincere@april.biz",
		Address{
			"Kulas Light",
			"Apt. 556",
			"Gwenborough",
			"92998-3874",
			Geo{
				"-37.3159",
				"81.1496",
			},
		},
		"1-770-736-8031 x56442",
		"hildegard.org",
		Company{
			"Romaguera-Crona",
			"Multi-layered client-server neural-net",
			"harness real-time e-markets",
		},
	}}

func GetUserFromJsonPlaceholder() {
	response, err := http.Get(jsonPlaceholderEndpoint)
	lfunc.CheckNilErr(err)
	defer response.Body.Close()
	bytes, err := io.ReadAll(response.Body)
	data := string(bytes)
	if err != nil {
		panic(err)
	}
	lfunc.CheckNilErr(err)
	print(data)
}

func MimicJsonPlaceholder() {
	response, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		panic(err)
	}
	print(string(response))
}
