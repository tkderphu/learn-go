package main
import "fmt"
func main() {
	
	var x int  = 5
	var ptr1 *int = &x
	var ptr2 **int = &ptr1
	

	fmt.Println("value of x: ", x)
	fmt.Println("Address of x: ", ptr1)
	fmt.Println("Value of ptr1: ", ptr1)
	fmt.Println("Address of ptr1: ", ptr2)

	// *ptr2 = new(int)
	// *ptr1 = 55555

	fmt.Println("Value after change: ",x)
	fmt.Println("Value ptr1", *ptr1)

}