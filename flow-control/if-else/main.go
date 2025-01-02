package main

import "fmt"

func main() {

	age := 16
	if age >= 18 {
		fmt.Println("Ban du ket hon")
	} else if age >= 16 {
		fmt.Println("ban dang hoc cap 3")
	} else {
		fmt.Println("ban chua du tuoi ket hon")
	}

	//we can also compact to
	if age := 18; age >= 18 {
		fmt.Println("ban da tot nghiep cap 3")
	}

	season := "Mua he"
	switch season {
	case "Mua he":
		fmt.Println("Mua he roi")
		fallthrough //-> chuyen toi case tiep theo va thuc hien no
	case "Mua xuan":
		fmt.Println("Mua xuan roi")
	case "Mua dong":
		fmt.Println("MUA DONG")
	default:
		fmt.Println("Mua thu")
	}

}
