package dytt

import (
	"crawler/engine"
	"log"
	"regexp"
)

var (
	//<a href='/html/gndy/dyzz/20200424/59958.html'>2019年高分获奖剧情《看不见的女人》BD中英</a><br>
	movieListRe = regexp.MustCompile(`<a href='(/html/gndy/dyzz/[\d]+/[^']+.html)'>`)
)

func ParseMovieList(contents []byte, url string) engine.ParseResult {
	all := movieListRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range all {
		result.Requests = append(result.Requests, engine.Request{
			Url:    "https://www.dytt8.net" + string(m[1]),
			Parser: engine.NewFuncParser(ParseMovie, "ParseMovie"),
		})
	}
	log.Printf("ParseMovieList found: %d\n", len(all))
	return result
}
