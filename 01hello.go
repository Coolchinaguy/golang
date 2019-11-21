// 01hello.go
package main

import "fmt"

func main() {
/*	var x []int
	fmt.Println(len(strconv.Itoa(x)))
	for k,v := range x {
		fmt.Println(k,v)
	}*/
	s := "A"
	for i,v := range s {
		//strings := byte(v)
		fmt.Printf("i = %d,v = %v\n",i,v)
	}
	//isPalindrome(121)
}

/*func isPalindrome(x int) int {
	if x < 10 {
		return x
	}
	y := 0
	for x > 0 {
		y *= 10
		y += x % 10
		x /= 10
	}
	if x == y {
		fmt.Println(y)
		return y
	} else {
		fmt.Println(x,y)
		fmt.Println("0")
		return 0
	}
}*/
