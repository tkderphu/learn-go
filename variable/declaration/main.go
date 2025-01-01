package main
import "fmt"

func main() {
	// var foo string //declare but is not assigned value
	var test string = "hello world"

	//multiple declarations
	// var foo, bar string = "test1", "test2"
	var (
		foo string = "test1"
		bar string = "test2"
	)

	//system will be inferred
	var my_type = "hello world" //->type: string
	println(my_type)

	//shorthand declaration, this is only be used in func body
	shortHand := "short hand"
	println(shortHand)


	//constant
	const constant = "this is a constant"
	println(constant)
	

	fmt.Println(foo + bar)
	fmt.Println(test)
	println("-----------------------------------this is example----------")
	println("-----------Ex: Enter your info and show on console-----------")

	var (
		fisrt_name string = "phu"
		last_name string = "quang"
	)
	age, address := 20, "Bac ninh"
	var major, student_id = "CNTT", "B22DCCN621"

	fmt.Printf("FirstName: %s\nLastName: %s\nAge: %d\nAddress: %s\nMajor: %s\nStudentId: %s",
	fisrt_name, last_name, age, address, major, student_id)	


}