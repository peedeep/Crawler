package engine

import (
	"crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("requests size %d, Fetching %s", len(requests), r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
			continue
		}
		parser := r.Parser
		parseResult := parser.Parse(body, r.Url)
		requests = append(requests, parseResult.Requests...)

		//for _, item := range parseResult.Items {
		//	log.Printf("Got item %v", item)
		//}
	}
}
