package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//option one to declare var
	var name string = "wael"
	//option two to declare var
	var age = 27
	//option three to declare var
	isMerged := true
	// &anyVar mean the addreace of var on memory
	pointerToAge := &age
	var newPointer *int // to init new pointer

	fmt.Println("the value of pointer is ,", newPointer) //nil
	fmt.Println(age)
	// *anyVar mean the value of the pointer location

	*pointerToAge = *pointerToAge + 1
	newPointer = pointerToAge
	fmt.Println("the value of pointer is ,", *newPointer)

	fmt.Println(age)

	fmt.Println(isMerged)
	fmt.Println("Hi my name " + name + " , my age  " + ", is marege ")

	// to read data from buffer momery via standard input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter company name \n")

	// reader return meltable values so we add comma ,
	// the second value the error so if you need to ignore it use _
	companyName, _ := reader.ReadString('\n')
	fmt.Println("I will generate cv for ", companyName)
	fmt.Printf("the ttpe of input %T \n", companyName)

	fmt.Println("How many eyers of experiences you have ")

	inputExperiences, _ := reader.ReadString('\n')
	experence, err := strconv.ParseInt(strings.TrimSpace(inputExperiences), 16, 64)
	checkNilError(err)
	genreatedTime := time.Now()
	fmt.Println("CV generated on ", genreatedTime.Format("2006/01/02 Monday 01:02:04"))
	fmt.Println("your experence ", experence+1)

	// array
	var brothers = [3]string{"yazan ", "wael", "ahmad"}
	fmt.Println("the brothers are", brothers)
	var sisters [1]string
	sisters[0] = "rania"
	fmt.Println("sisters count is ", len(sisters), " and they are ", sisters)

	for index, member := range brothers {
		fmt.Println("brother name ", member, " at index ", index)
	}
	for member := range brothers {
		fmt.Println("brother name ", " at index ", brothers[member])
		if brothers[member] == "wael" {
			goto loc
		}
	}
	// goto statment label
loc:
	fmt.Println("print numbers from 1 to 5")

	// map
	// var family map[string]string
	family := make(map[string]string)
	family["father"] = "Khaleel"
	family["mother"] = "Suha"
	fmt.Println("family members ", family)

	p := Person{name: "wael", age: 27}
	fmt.Println("person info ", p)
	p.setAge(30)
	fmt.Println("person age ", p.getAge())
	p.getPersonInfo()

	createCVFile("wael", 27, "Google")
}

func isFileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
func readCVFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("error opening file ", err)
		return
	}
	content := string(dataByte)
	fmt.Println("the cv content is ", content)

}

func createCVFile(name string, age int, company string) {
	isFileCreated := isFileExists("cv.txt")
	if isFileCreated {
		readCVFile("cv.txt")
		return
	}

	file, err := os.Create("cv.txt")
	if err != nil {
		fmt.Println("error creating file ", err)
		return
	}
	io.WriteString(file, "the cv of "+name)
	defer file.Close()

	defer file.Close()
}

// function to check nil error
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

type Person struct {
	name string
	age  int
}

// method to get age
func (p Person) getAge() int {
	return p.age
}

// func to set age
func (p *Person) setAge(age int) {
	p.age = age
}

// method to get person info

func (p Person) getPersonInfo() {
	// example of defer keyword , defer used to delay the execution of a function until the surrounding function returns
	defer fmt.Println("get person info called")
	fmt.Sprintf("name %s , age %d", p.name, p.age)
}
