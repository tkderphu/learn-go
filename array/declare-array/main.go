package main
import "fmt"
func main() {
	var a [5] int

	fmt.Println(a[1])

	a[1] = 5

	fmt.Println(a)



	fmt.Println(a[1])
}