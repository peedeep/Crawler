package main

import (
	"crawler/distribute/config"
	"crawler/distribute/persist"
	"crawler/distribute/rpcsupport"
	"flag"
	"fmt"
	"github.com/olivere/elastic.v7"
	"log"
)

var port = flag.Int("port", 1234, "The port for itemsaver server to listen on.")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("Port must be specified")
		return
	}
	log.Printf("Itemsaver server listening on port %d.", *port)

	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
}

func serveRpc(host string, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
