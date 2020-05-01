package world

import "fmt"

// Hello function prints greeting messgae and returns name
func Hello(name string) string {
	fmt.Println(fmt.Sprintf("Hello, %s", name))
	return name
}

// ReturnTest returns function's value
func ReturnTest() int {
	return returnOne()
}

func returnOne() int {
	return 1
}
