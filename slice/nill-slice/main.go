package main
import "fmt"
func main() {
	var s [] string
	var arr = make([]string, 5)
	fmt.Println(s)
	fmt.Println(arr)
	s = append(s, "ss", "ds", "fsd", "cd")
	fmt.Println(s)




	test := append(arr, "f", "f")
	fmt.Println(test)
}