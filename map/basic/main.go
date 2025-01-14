package main
import "fmt"
func main() {
	//declare map :=> var x = map[keyType]ValueType{
    // key1: value1,key2: value2
	//}
		var test = make(map[int]string)

		if test == nil {
			fmt.Println("test is nil")
		} else {
			fmt.Println("test is not nil")
		}

		test[100] = "one hundred"

		fmt.Println(test)

		var b map[int]string
		fmt.Println(b)
		if (b == nil) {
			fmt.Println("nill")
		}
		var myMap = map[string]string{}
		fmt.Println(myMap)
		fmt.Println(myMap["tran nam"])
		
		var x [][]int

		fmt.Println(x)

}