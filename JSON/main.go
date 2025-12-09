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

func (p Person) getPersonInfo() string {
	fmt.Println("person info method called")
	return p.Name
}

func (p *Person) setPersonAge(Age int) {
	p.Age = Age

}

/*
	In Go, an interface is a contract that defines a set of methods without providing any implementation.

	Any struct that implements these methods automatically satisfies the interface, without the need for keywords like implements or extends (unlike Java or PHP).

	This design enables powerful polymorphism, allowing you to write flexible, behavior-driven code that depends on what a type can do, not what it is.
*/

type Personalization interface {
	getPersonInfo() string // the method name ,it's parameters  and what it returns
	setPersonAge(Age int)
}

// so any struct that has implemented the methods of the interface Personalization
// can be passed to this function

func getName(p Personalization) {
	fmt.Println("called by interface")
	fmt.Println(p.getPersonInfo())
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
