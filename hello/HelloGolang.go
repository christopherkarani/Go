package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Post
type Post struct {
	UserID int
	ID     int
	Title  string
	Body   string
}

func retrieveJSONPlaceHolderJSON() {
	test := "https://jsonplaceholder.typicode.com/posts/1"
	rs, err := http.Get(test)
	if err != nil {
		panic(err)
	}
	defer rs.Body.Close()
	bodyBytes, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		panic(err)
	}

	var post Post
	postJSON := string(bodyBytes)
	json.Unmarshal([]byte(postJSON), &post)

	returnString := fmt.Sprintf("Post Description: userID: %d, ID: %d, title: %s, body: %s", post.UserID, post.ID, post.Title, post.Body)
	println(returnString)
}

func main() {

	name, age := greetKiannah()

	println(name, age)

}

func getPokeAPIJOSN() {
	baseURL := "http://pokeapi.co/api/v2/pokemon/1"

	response, error := http.Get(baseURL)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	data, error := ioutil.ReadAll(response.Body)
	if error != nil {
		panic(error)
	}

	responseString := string(data)

	println(responseString)

}

func greeting() {
	greetingString := "Hello Kiannah"
	fmt.Println(greetingString)
}

func greetingsWithName(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

func greetKiannah() (string, int) {
	return "Kiannah", 38
}
