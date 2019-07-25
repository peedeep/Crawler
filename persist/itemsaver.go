package persist

import (
	"context"
	"crawler/engine"
	"github.com/olivere/elastic"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item saver: got item #%d: %v", itemCount, item)
			itemCount++
			err := save(client, item, index)
			if err != nil {
				log.Printf("Item saver: error saveing item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}

func save(client *elastic.Client, item engine.Item, index string) error {
	_, err := client.Index().
		Index(index).
		Type(item.Type).
		Id(item.Id).
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}
