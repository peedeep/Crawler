package persist

import (
	"context"
	"crawler/engine"
	"crawler/model"
	"encoding/json"
	"github.com/olivere/elastic.v7"
	"testing"
)

func TestItemSaver(t *testing.T) {
	expected := engine.Item{
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
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	
	err = save(client, expected, "dating_test")
	if err != nil {
		panic(err)
	}

	// TODO: Try to start up elastic search
	// here using docker go client
	result, err := client.Get().
		Index("dating_test").
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	//t.Logf("%+v\n", result)
	t.Logf("%s", result.Source)

	var actual engine.Item
	err = json.Unmarshal(result.Source, &actual)
	if err != nil {
		panic(err)
	}

	actualProfile, _ := model.FromJsonOjb(actual.Payload)
	actual.Payload = actualProfile

	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
