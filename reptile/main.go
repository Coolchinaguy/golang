package main

import (
	"reptile/58tongchen/parser"
	"reptile/engine"
)

func main(){
	engine.Run(engine.Request{
		Url:"https://www.58.com/job/",
		ParserFunc:parser.ParseCityList,
	})
}




