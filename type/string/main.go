package main

import (
	"fmt"
	"strings"
)
func main() {

	myFullName := "NgUyen Quang PhU"
	myArrays := strings.Split(myFullName, "")
	fmt.Println(myArrays)
	fmt.Printf("myFullName: %s\n", myFullName)


	var str string = `hello guys
		xin chao
		please????
	`

	println(str)

}