package main

import (
	"crawler/engine"
	"crawler/parser/dytt"
	"crawler/persist"
	"crawler/scheduler"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_movie")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueueScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}
	//e.Run(engine.Request{
	//	Url:    "https://www.zhenai.com/zhenghun",
	//	Parser: parser.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	//})
	e.Run(engine.Request{
		Url:    "https://www.dytt8.net/",
		Parser: engine.NewFuncParser(dytt.ParseMovieList, "ParseMovieList"),
	})

}
