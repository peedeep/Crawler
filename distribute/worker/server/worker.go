package main

import (
	"crawler/distribute/config"
	"crawler/distribute/rpcsupport"
	"crawler/distribute/worker"
	"fmt"
	"log"
)

func main() {
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", config.WorkerPort0), worker.CrawlService{}))
}
