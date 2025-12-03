package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	EncodeToJSON()
	DecodeFromJSON()

}

type Person struct {
	// all  field should  be  started with capital to be exported
	Name        string   `json:"MEMBER_NAME"`             // `json:"MEMBER_NAME"` Called struct tag
	Age         int      `json:"-"`                       // "-" to ignore this field during encoding
	Nationality []string `json:"nationalities,omitempty"` // "omitempty" to ignore this field if it's empty
}

func EncodeToJSON() {
	jsonData := []Person{
		{"wael", 27, []string{"Egypt", "KSA"}},
		{"ali", 30, []string{"UAE", "KSA"}},
		{"omar", 25, nil},
	}
	jsonByte, _ := json.MarshalIndent(jsonData, "", " ")

	fmt.Printf("the json data is %s \n ", string(jsonByte))

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
				fmt.Printf("the person %s is found with age %d \n", person.Name, person.Age)
			}
		}
	}

}
