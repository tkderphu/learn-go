package main

import (
	"fmt"
	"text/template/parse"
)
func main() {
	i := 5
	gpa := 5.5
	cpa := uint(gpa) //conversion gpa -> cpa(int)

	float32I := float32(i)


	fmt.Printf("%T %T %T, %T", i, gpa, cpa, float32I)

}