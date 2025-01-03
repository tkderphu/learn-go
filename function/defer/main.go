//defer is a keywork for postpones execution of function
//until the surrounding function returns
//Defer is incredibly useful and is commonly 
//used for doing cleanup or error handling
package main
import "fmt"
func main() {
	str := "Nguyen quang phu"

	defer hello_world()
	fmt.Println(str)
}

func hello_world() {
	fmt.Println("Xin chao moi nguoi")
}