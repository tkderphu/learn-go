package main
import "fmt"
func main() {
	
	var x int = 5
	var mem *int = &x
	fmt.Println("Value of x: ", x)
	fmt.Println("Address in memory of x: ", mem)


	*mem = 1

	fmt.Println("Value of x after change: ", x)


	var newPointer *int

	fmt.Println(newPointer)



	ptr := new(int)

	fmt.Println(ptr)
	fmt.Println("Value ptr: ", *ptr)
	*ptr = 55
	fmt.Printf("Ptr = %#x, Ptr value = %d\n", ptr, *ptr)
}