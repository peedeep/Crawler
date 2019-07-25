package parser

import (
	"crawler/engine"
	"log"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityRe    = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte, url string) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		//log.Printf("parse city user name: %s", name)
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: NewProfileParser(name),
		})
	}
	matches = cityRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: NewFuncParser(ParseCity, "ParseCity"),
		})
	}
	log.Printf("parse city result size: %d", len(result.Requests))
	return result
}
