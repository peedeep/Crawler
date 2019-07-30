package client

import (
	"crawler/distribute/config"
	"crawler/distribute/rpcsupport"
	"crawler/distribute/worker"
	"crawler/engine"
	"fmt"
	"net/rpc"
)

func CreateProcessor() (engine.Processor, error) {
	client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.WorkerPort0))
	if err != nil {
		return nil, err
	}
	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(req)
		var sResult engine.SerializedParseResult
		// TODO call rpc to craw data
		err := client.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult)
	}, nil
}

func CreateProcessors(clientChan chan *rpc.Client) (engine.Processor, error) {
	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(req)
		var sResult engine.SerializedParseResult
		// TODO call rpc to craw data
		client := <-clientChan
		err := client.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult)
	}, nil
}
