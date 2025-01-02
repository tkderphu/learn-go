package main
import "fmt"

type MyAlias = string
func main() {


	var myName MyAlias = "nguyen quang phu"


	fmt.Printf("%T\n", myName) 

	fmt.Print(myName)

}