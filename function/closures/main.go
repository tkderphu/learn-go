package main
import "fmt"
func main() {
	add := myFunction()
	add(5) //sum = 5
	res:=add(6) // sum = 11
	fmt.Println(res)
}

func myFunction() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}