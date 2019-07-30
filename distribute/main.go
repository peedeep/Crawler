package main

import (
	"crawler/distribute/config"
	"crawler/distribute/persist/client"
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
	"fmt"
)

func main() {
	itemChan, err := client.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
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
