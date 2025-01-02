package main
import "fmt"
func main() {
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}
		if i == 4 {
			break
		}
		fmt.Println(i)
	}

	i := 0
	for ; i <= 5; {
		i++
	}

	//loop forever
	for {

	}
}