package main
import "fmt"
func main() {
	var a = [5] int {1, 2, 3, 4, 5}

	b := [5]int{2}

	c := [...] int {2, 3, 4, 5}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}