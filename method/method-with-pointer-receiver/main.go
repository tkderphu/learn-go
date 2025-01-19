package main

import "fmt"


type Person struct {
	name string;
	age int;
}


func (person* Person) setAge(age int) {
	person.age = age
}
func (person* Person) setName(name string) {
	person.name = name
}

func main() {


	person := Person {
		name: "Nguyen quang phu",
		age: 50,
	}
	fmt.Println("Before set value for person object: ", person)
	person.setName("Kitoshima")
	fmt.Println("After set value for person name: ", person)
	person.setAge(15)
	fmt.Println("After set value for person age: ", person)

}