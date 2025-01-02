package main
import "fmt"

type MyDefine string

func main() {
	var myName MyDefine = "Nguyen quang phu"
	fmt.Printf("%T", myName)

	/*
		What different between alias and define
		When using define => you define a new type
		else => you only assign type for alias type
		Example:
			type alias = string
			type def string

			var myname alias = "quang phu"
			var assign string = alias => ok

			var myName def = "string"
			var assign string = def => error (two types are diffent)

	*/

	fmt.Print("\n", myName)
}