package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
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
	performGetRequest(myurl)

}

func performGetRequest(reqestedUrl string) {
	response, err := http.Get(reqestedUrl)
	checkNilError(err)
	// caller responsibility to close the body
	defer response.Body.Close()
	responseByte, err := io.ReadAll(response.Body)
	checkNilError(err)
	// option one to read the response body
	fmt.Println("Response Body:", string(responseByte))
	// option two to read the response body
	var responseBody strings.Builder
	responseBody.Write(responseByte)
	fmt.Println("Response Body:", responseBody.String())

	// other information about response
	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Content Length:", response.ContentLength)
}

func performPostRequest(requestedUrl string) {
	requestBody := strings.NewReader(`
		{
			"name":"wael",
			"age":27
		}
	`)
	response, err := http.Post(requestedUrl, "application/json", requestBody)
	checkNilError(err)
	defer response.Body.Close()
}

func performPostFormRequest(requestedUrl string) {
	formData := url.Values{}
	formData.Add("name", "wael")
	formData.Add("age", "27")
	response, err := http.PostForm(requestedUrl, formData)
	checkNilError(err)
	defer response.Body.Close()
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}

}
