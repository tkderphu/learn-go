package main
import "fmt"
func main() {
	myMap := map[int]int {
		1:1, 2:4,
	}
	fmt.Println(myMap)

	delete(myMap, 2)
	fmt.Println(myMap)
}