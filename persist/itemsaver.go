package persist

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"log"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item saver: got item #%d: %v", itemCount, item)
			itemCount++
			_, err := save(item)
			if err != nil {
				log.Printf("Item saver: error saveing item %v: %v", item, err)
			}
		}
	}()
	return out
}

func save(item interface{}) (string, error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return "", nil
	}
	resp, err := client.Index().
		Index("dating_profile").
		Type("zhenai").
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		return "", nil
	}

	fmt.Println(resp)
	return resp.Id, nil
}
