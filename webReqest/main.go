package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const myurl string = "https://jsonplaceholder.typicode.com:443/todos/1?name=wael&age=27"

func main() {
	// parsing URL
	result, _ := url.Parse(myurl)
	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("Path:", result.Path)
	fmt.Println("RawQuery:", result.RawQuery)
	fmt.Println("Host", result.Port())
	fmt.Println("Query Params:", result.Query()) // the datatype is map[string][]string

	for key, val := range result.Query() {
		fmt.Printf("key is %v and value is %v \n", key, val)
	}

	fmt.Println("name param:", result.Query()["name"])
	fmt.Println("age param:", result.Query()["age"])
	response, err := http.Get(myurl)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("Response Type  %T \n", response)
		response, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err.Error())
		} else {
			fmt.Println(string(response))
		}
	}
	// caller responsibility to close the body
	defer response.Body.Close()

}
