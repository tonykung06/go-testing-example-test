package pack

import (
	"fmt"
)

//package level example
func Example() {
	result := FunctionToBeTest(1, 2)

	fmt.Println(result)

	// Output:
	// 3
}

//package function level example
func ExampleFunctionToBeTest() {
	result := FunctionToBeTest(1, 2)

	fmt.Println(result)

	// Output:
	// 3
}

func ExampleMyStruct_Method1() {
	myStruct := MyStruct{}
	result := myStruct.Method1()

	// Output:
	// 1
}
