package main

import "fmt"

/**
	with non-struct type  you must redefine type that you want to create a method
	because all of the non-struct type is not the same package with the method
	which you want creation
	
*/

type number int

func (num number) power(x number) number {
	var i number = 1
	var res number = 1
	for ;i <= x; i++{
		res *= num
	} 
	return res
}

func main() {
	var x number = 5
	fmt.Println(x.power(2))
}