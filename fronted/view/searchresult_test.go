package view

import (
	"crawler/engine"
	"crawler/fronted/model"
	common "crawler/model"
	"os"
	"testing"
)

func TestCreateSearchResultView(t *testing.T) {

	view := CreateSearchResultView("template.html")

	out, err := os.Create("template_test.html")

	page := model.SearchResult{}
	page.Hits = 123
	profile := common.Profile{
		"name",
		"name",
		"name",
		"name",
		111,
		"name",
		"name",
		"name",
		"name",
		"name",
		"name",
		"name",
		"name",
	}
	item := engine.Item{
		"http://album.zhenai.com/u/108906739",
		"type",
		"1",
		profile,
	}
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}
	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}
}
