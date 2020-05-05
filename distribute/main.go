package main

import (
	itemsaver "crawler/distribute/persist/client"
	"crawler/distribute/rpcsupport"
	worker "crawler/distribute/worker/client"
	parser2 "crawler/dytt/parser"
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String("itemSaverHost", ":1234", "item saver host")
	workerHosts   = flag.String("workerHosts", ":9000", "worker hosts(comma separated)")
)

func main() {
	flag.Parse()
	if *itemSaverHost == "" || *workerHosts == "" {
		fmt.Println("Port must be specified.")
		return
	}

	// fmt.Sprintf(":%d", config.ItemSaverPort)
	itemChan, err := itemsaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	//processor, err := worker.CreateProcessor()
	pool := createClientPool(strings.Split(*workerHosts, ","))
	processor, err := worker.CreateProcessors(pool)
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
		Url:    "https://www.dytt8.net/",
		Parser: parser.NewFuncParser(parser2.ParseMovieList, "ParseMovieList"),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err != nil {
			log.Printf("Error connecting to %s: %v", h, client)
		} else {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		}
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
