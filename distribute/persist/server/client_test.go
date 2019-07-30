package main

import (
	"crawler/distribute/config"
	"crawler/distribute/rpcsupport"
	"crawler/engine"
	"crawler/model"
	"fmt"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/81397582",
		Type: "zhenai",
		Id:   "81397582",
		Payload: model.Profile{
			Name:     "安静的雪",
			Age:      "34",
			Gender:   "女",
			Marriage: "离异",
		},
	}

	const host = ":1234"
	go serveRpc(host, "test1")
	time.Sleep(2 * time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	var result = ""
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	} else {
		fmt.Println(result)
	}
}
