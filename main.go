package main

import (
	"crawler/engine"
	"crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:    "https://www.zhenai.com/zhenghun",
		Parser: parser.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})

}
