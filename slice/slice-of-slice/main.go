package main
import "fmt"
func main() {
	mySlice := [][]int {
		{1, 2, 3},
		{3, 5},
	}
	res := mySlice[1]
	res = append(mySlice[0], 4)
	fmt.Println(res)

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}

	for index, _ := range res {
		fmt.Println(index)
	}

}