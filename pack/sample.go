package pack

type MyStruct struct {
}

func (ms *MyStruct) Method1() {
	return 1
}

func FunctionToBeTest(num1, num2 int) int {
	return num1 + num2
}
