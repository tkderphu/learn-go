package main
import "fmt"
func main() {
	var c1 complex64 = complex(2, 5)
	var c2 complex64 = 2-5i

	fmt.Print(c1 * c2)
}