package main
import "fmt"
func main() {

	isFalse := "queue" == "queue" && 5 > 50
	isTrue := 5 >= 7 || "q"=="q"
	fmt.Print(isTrue, "\n")
	fmt.Print(isFalse, "\n")

	fmt.Println("Negative:", !isFalse)
}