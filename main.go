package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}
	//e.Run(engine.Request{
	//	Url:    "https://www.zhenai.com/zhenghun",
	//	Parser: parser.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	//})
	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun/shanghai",
		Parser: parser.NewFuncParser(parser.ParseCity, "ParseCity"),
	})

}
