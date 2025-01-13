package main
import "fmt"
func main() {
	/**
		Slice is similarly arr but its has many func for flexibility and power than arr

	**/

	//declaration
	mySlice := []int {1, 2, 3}

	fmt.Println(mySlice)
	//number of element arr can put into
	fmt.Println(cap(mySlice))
	//number of element current of arr
	fmt.Println(len(mySlice))
	
	//create slice from arr
	myArr := [6]int {1, 2, 3, 4, 5, 6}
	mySlice = myArr[1:5]
	fmt.Println(mySlice)
	fmt.Println(cap(mySlice))
	fmt.Println(len(mySlice))


	//create slice with make function
	mySlice = make([]int, 4)
	fmt.Println(mySlice)
	res := append(mySlice, 1,4,5,6,7,8,9,5,3)
	fmt.Println(res) 


	otherSlice := []int {-1, -2, -3, -4}

	// '...': mean get all elements from slice
	res = append(res, otherSlice...)

	fmt.Println(res)


	//copy function: for get element from target slice
	target := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	copySlice := make([]int, 5)
	copy(copySlice, target)
	fmt.Println(copySlice)




	


}