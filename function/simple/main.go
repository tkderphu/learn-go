package main
import "fmt"


func main() {
	hello_world_2("go tutorial", "test")

	sum := sumTwoInteger(5, 5)
	fmt.Printf("\n%d", sum)

	/**
	multiple value return from function
	*/
	s, i := multipleValue("hello guys")
	fmt.Printf("\n%s %d\n", s,i)


	v, k := myFuntion("hehehe")

	fmt.Println(v, k)


	//function is value
	fn := func () string  {
		fmt.Println("function is value")
		return "hello funciton is vlaue"
	}
	test := fn()
	fmt.Println(test)


}

func hello_world() {
	fmt.Println("simple function")
}

func hello_world_1(message, name string) {
	fmt.Print("Hello world: " + message)
}
func hello_world_2(message string, name string) {
	fmt.Print("Hello world: " + message + " - - - " + name)
}

func sumTwoInteger(a int, b int) int{
	return a + b
}

func multipleValue(p1 string) (string, int) {
	msg := fmt.Sprintf("%s function", p1)
	return msg, 10
}

//name return
func myFuntion(message string) (s string, i int) {
	s = fmt.Sprintf("%s function", message)
	i = 10
	return
}