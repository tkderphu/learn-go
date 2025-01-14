package main
import ("fmt")
func main() {
	myMap := make(map[string]int)

	myMap["test"] = 3433
	myMap["tetua"] = 3646
	myMap["messi"] = 8
	myMap["yamal"] = 19


	// name: value, ok: bool -> true key exists else false
	name, ok := myMap["ve"]

	fmt.Println(name, ok)

	if !isKeyExists(myMap, "3434") {
		fmt.Println("key is not exists")
	}

}

func isKeyExists(yourMap map[string]int, key string) bool {
	_, ok := yourMap[key]

	return ok
}

