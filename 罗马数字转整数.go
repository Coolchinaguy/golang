package main

import (
	"fmt"
	"strings")
/*
func judgeRomanTwo(s string) int{
	judes := map[string]int{"I":1,"V":5,"X":10,"L":50,"C":100,"D":500,"M":1000}
	return judes
}*/

func main() {
	s := "IV"
	jss := 0
	for _,v := range s {
		b := string(v)
		js := judgeRoman(b)
		jss += js
	}
	if strings.Contains(s,"CD") || strings.Contains(s,"CM"){
		jss -= 200
	}
	if strings.Contains(s,"IV") || strings.Contains(s,"IX"){
		jss -= 2
	}
	if strings.Contains(s,"XL") || strings.Contains(s,"XC"){
		jss -= 20
	}
	fmt.Println(jss)
}


func judgeRoman(s string) int{
	if s == "I" {
		return 1
	}
	if s == "V" {
		return 5
	}
	if s == "X" {
		return 10
	}
	if s == "L" {
		return 50
	}
	if s == "C" {
		return 100
	}
	if s == "D" {
		return 500
	}
	if s == "M" {
		return 1000
	}
	return 0
}