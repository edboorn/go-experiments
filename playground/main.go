package main

import "fmt"

// Can shortern variable declerations if both parameters are the same type
func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func main() {
	fmt.Println(add(1, 2))
}

func functionTest(data1 string, data2 string, data3 int) {
	data1 = data1

}
