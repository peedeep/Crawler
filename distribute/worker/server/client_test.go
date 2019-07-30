package main

import (
	"crawler/distribute/config"
	"crawler/distribute/rpcsupport"
	"crawler/distribute/worker"
	"crawler/engine"
	"fmt"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(2 * time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := engine.SerializedRequest{
		Url: "http://album.zhenai.com/u/81397582",
		Parser: engine.SerializedParser{
			Name: config.ParseProfile,
			Args: "安静的雪",
		},
	}
	var result engine.SerializedParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Errorf("err: %v", err)
	} else {
		fmt.Println(result)
	}

}
