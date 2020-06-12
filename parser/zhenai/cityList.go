package zhenai

import (
	"crawler/distribute/config"
	"crawler/engine"
	"log"
	"regexp"
)

//<a href="http://www.zhenai.com/zhenghun/akesu" data-v-5e16505f="">阿克苏</a>
const cityListRegex = `<a href="(http://www.zhenai.com/zhenghun/[a-zA-z0-9]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	compile := regexp.MustCompile(cityListRegex)
	all := compile.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range all {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
		})
	}
	log.Printf("Matches found: %d\n", len(all))
	return result
}
