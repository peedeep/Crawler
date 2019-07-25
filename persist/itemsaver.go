package persist

import (
	"context"
	"crawler/engine"
	"github.com/olivere/elastic"
	"log"
)

func ItemSaver() chan engine.Item {
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item saver: got item #%d: %v", itemCount, item)
			itemCount++
			err := save(item)
			if err != nil {
				log.Printf("Item saver: error saveing item %v: %v", item, err)
			}
		}
	}()
	return out
}

func save(item engine.Item) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return err
	}
	_, err = client.Index().
		Index("dating_profile").
		Type(item.Type).
		Id(item.Id).
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}
