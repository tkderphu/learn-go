package main
import "fmt"
func main() {
	myMap := map[string]int {
		"phu": 20,
		"phong": 22,
		"tran": 35,
	}

	for name, age := range myMap {
		fmt.Println(name, age)
	}
}