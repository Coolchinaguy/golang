package main

import "fmt"

func main() {
	const (
		a    = iota
		b    = iota
		c, d = iota, iota
	)
	fmt.Printf("a =%d,b =%d,c =%d,d =%d\n", a, b, c, d)

	const (
		j1 = iota
		j2
		j3
	)

	fmt.Printf("j1 = %d,j2 = %d,j3 = %d", j1, j2, j3)
}
