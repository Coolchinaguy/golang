package parser

import (
	"regexp"
	"reptile/engine"
)

func ParseCityList(contents []byte)engine.ParseResult  {
	re := regexp.MustCompile(
		`<a target="_blank" href="(//[a-z]{2}.58.com/job.shtml)">([^<]{2})</a>`)
	matchs := re.FindAllSubmatch(contents,-1)
	result := engine.ParseResult{}
	for _,m := range matchs{
		result.Items = append(result.Items,m[2])
		result.Requests = append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParserFunc:engine.NilParser,
		})
	}
	return  result
}

