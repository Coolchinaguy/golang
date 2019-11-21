package main

import "fmt"

func main() {
	a := 10.5 //自动推导类型，包括创建变量，并对变量赋值
	fmt.Println("a=", a)
	fmt.Printf("a the type of a is %T\n", a)

	var b byte = 'a'
	fmt.Printf("b=%c", b)
}
