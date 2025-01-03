//init function is that will be executed before
//main function
//is usually use for connection database or setting enviroment
package main
import "fmt"

func init() {
	fmt.Println("init function called 1")
}
func init() {
	fmt.Println("init function call2")
}

func main() {
	fmt.Println("main function called")
}