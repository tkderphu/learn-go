package main
import "fmt"

type Person struct {
	name string;
	age int;
	address string;
	sex string;
}

type Customer struct {
	person Person;
	username, password string

}

func main() {

	var person Person = Person{
		name: "Quang Phu",
		address: "Yen Phong, Bac Ninh",
		age: 20,
	}
	fmt.Println(person)
	fmt.Println(person.address)


	customer := Customer{
		person: person,
		username: "hello",
		password: "hello",
	}

	fmt.Println(customer)




	ptrCustomer := new(Customer)
	ptrCustomer = &Customer{
		username: "test",
	}
	fmt.Println(ptrCustomer)
	fmt.Println(ptrCustomer.username)

}