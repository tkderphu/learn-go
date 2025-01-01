package main
import "fmt"
func main() {
	var (
		age int32 //->0
		name string // -> ""
		sex bool //-> false
		gpa float32 // -> 0.0000
	)
	fmt.Printf("%d %s %v %f", age, name, sex, gpa)
}