package persist

import (
	"context"
	"crawler/model"
	"encoding/json"
	"github.com/olivere/elastic"
	"testing"
)

func TestItemSaver(t *testing.T) {
	expected := model.Profile{
		Name:     "安静的雪",
		Age:      "34",
		Gender:   "女",
		Marriage: "离异",
	}
	id, err := save(expected)
	if err != nil {
		panic(err)
	}
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	// TODO: Try to start up elastic search
	// here using docker go client
	result, err := client.Get().Index("dating_profile").
		Type("zhenai").
		Id(id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	//t.Logf("%+v\n", result)
	t.Logf("%s", result.Source)

	var actual model.Profile
	err = json.Unmarshal(result.Source, &actual)
	if err != nil {
		panic(err)
	}
	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
