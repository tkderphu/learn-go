package main
import "fmt"
func main() {
	/**
	map is reference type when you assign var type map to
	another var type map and you change value so original map also is changed
	**/
	myMap := map[int]int {
		1:1, 2:5,
	}
	referenceType := myMap
	fmt.Println(myMap)
	referenceType[1] = 5

	fmt.Println(myMap)
}