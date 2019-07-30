package main

import (
	"crawler/distribute/config"
	"crawler/distribute/persist"
	"crawler/distribute/rpcsupport"
	"fmt"
	"github.com/olivere/elastic.v7"
	"log"
)

func main() {
	log.Fatal(serveRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ElasticIndex))
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
