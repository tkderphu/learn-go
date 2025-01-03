package main
import "fmt"
func main() {	
	sum := variadicFunc(1, 2, 3, 4, 5)
	fmt.Println(sum)
}

func variadicFunc(values...int) int{
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}
