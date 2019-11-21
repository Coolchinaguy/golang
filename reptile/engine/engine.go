package engine

import (
	"log"
	"reptile/fetcher"
)

func Run(seeds ...Request) {
	var requsets []Request
	for _,r := range seeds {
		requsets = append(requsets,r)
	}

	for len(requsets) > 0{
		r := requsets[0]
		requsets  = requsets[1:]

		body,err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Fatalf("Fetcher error "+"fetching url %s: %v",r.Url,err)
			continue
		}

		parseResult := r.ParserFunc(body)
		requsets = append(requsets,parseResult.Requests...)
		for _,item := range parseResult.Items {
			log.Printf("Got item %s", item)
		}
	}
}
