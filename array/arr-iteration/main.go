package main
import "fmt"
func main() {
	var x = [5] int{1, 2, 3, 4, 5}
	for i:= 0; i < len(x); i++ {
		fmt.Println(x[i])
	}
}