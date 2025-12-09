package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup // to wait for all goroutines to finish
var mu sync.Mutex     // to protect shared resources

var encoded int = 0
var channel = make(chan int, 5) // buffered channel to limit concurrency
func main() {

	for i := 0; i < 5; i++ {
		wg.Add(1) // increment the WaitGroup counter
		go EncodeToJSON()
		DecodeFromJSON()
		wg.Done() // decrement the WaitGroup counter
	}

	var x int
	for i := 0; i < 5; i++ {
		x = <-channel // release a slot in the buffered channel
		fmt.Printf("released channel slot %d \n", x)
	}
	// calledByInterface(&Person{"test", 0, nil})
	wg.Wait() // wait for all goroutines to finish
	time.Sleep(1 * time.Second)
	defer fmt.Printf("total encoded json calls %d \n", encoded)
}

type Person struct {
	// all  field should  be  started with capital to be exported
	Name        string   `json:"MEMBER_NAME"`             // `json:"MEMBER_NAME"` Called struct tag
	Age         int      `json:"-"`                       // "-" to ignore this field during encoding
	Nationality []string `json:"nationalities,omitempty"` // "omitempty" to ignore this field if it's empty
}

// the interface is used to define a contract for JSON operations
type JSONable interface {
	EncodeToJSON()
	DecodeFromJSON()
}

func calledByInterface(j JSONable) {
	fmt.Println("called by interface")
	j.EncodeToJSON()
	j.DecodeFromJSON()
}

func EncodeToJSON() {
	channel <- 1 // acquire a slot in the buffered channel

	mu.Lock()   // lock the mutex to protect shared resource
	encoded++   // this code will be executed by one goroutine at a time
	mu.Unlock() // unlock the mutex
	jsonData := []Person{
		{"wael", 27, []string{"Egypt", "KSA"}},
		{"ali", 30, []string{"UAE", "KSA"}},
		{"omar", 25, nil},
	}
	jsonByte, _ := json.MarshalIndent(jsonData, "", " ")

	fmt.Printf("***the json data is %s \n ", string(jsonByte))

}
func isValidJSON(jsonData []byte) bool {
	return json.Valid(jsonData)

}

func DecodeFromJSON() {
	jsonFromWeb := []byte(`[
		{
			"MEMBER_NAME": "wael",
			"Age": 27,
			"nationalities": [
			"Egypt",
			"KSA"
			]
 		},
		{
			"MEMBER_NAME": "ali",
			"Age": 30,
			"nationalities": [
			"UAE",
			"KSA"
			]
		}
	]`)
	if isValidJSON(jsonFromWeb) {
		var persons []Person
		json.Unmarshal(jsonFromWeb, &persons)
		fmt.Printf("the persons data is %v \n", persons)

		for _, person := range persons {
			if person.Name == "wael" {
				fmt.Printf("----the person %s is found with age %d \n", person.Name, person.Age)
			}
		}
	}

}
