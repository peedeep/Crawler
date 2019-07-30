package main

import (
	"crawler/distribute/config"
	itemsaver "crawler/distribute/persist/client"
	worker "crawler/distribute/worker/client"
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
	"fmt"
)

func main() {
	itemChan, err := itemsaver.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	processor, err := worker.CreateProcessor()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueueScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	//e.Run(engine.Request{
	//	Url:    "https://www.zhenai.com/zhenghun",
	//	Parser: parser.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	//})
	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun/nanchang",
		Parser: parser.NewFuncParser(parser.ParseCity, config.ParseCity),
	})

}
