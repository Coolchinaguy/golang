package main

import "fmt"

func main() {
	//numTable := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	//strings := "abcdefghijk"
	stringTable := make([][]string,1)
	stringTable[0][0] = "string"
	fmt.Println(cap(stringTable))
}
