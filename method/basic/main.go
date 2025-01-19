/**
	What is method? -> it is like function but its have some key feauture difference
	it must have receiver in its argument for accessing to properties of receiver,
	it must be created in the same package with the receiver
**/
/**
	Syntax: func (receiver Type) name_method(parameters...) return_type {

	}
*/
package main
import "fmt"

type Person struct {
	name string;
	age int;
}

func (person Person) display() {

	fmt.Println("Person age: ", person.age)
	fmt.Println("Person name: ", person.name)

	person.age = 50
}

func main() {
	person := Person{
		name: "Nguyen Quang Phu",
		age: 20,
	}

	person.display()
	
}