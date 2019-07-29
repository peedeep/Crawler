package controller

import (
	"context"
	"crawler/engine"
	"crawler/fronted/model"
	"crawler/fronted/view"
	"github.com/olivere/elastic.v7"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"regexp"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}

func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))

	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}

	page, err := h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

const pageSize = 10

func (h SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	log.Printf("get result %s, %d", q, from)
	var result model.SearchResult
	result.Query = q

	resp, err := h.client.
		Search("dating_profile").
		Type("zhenai").
		Query(elastic.NewQueryStringQuery(rewriteQueryString(q))).
		From(from).
		Do(context.Background())

	if err != nil {
		log.Printf("resp: %v\n, err: %v", resp, err)
		return result, err
	}
	log.Printf("result resp is: %v", resp)

	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)

	return result, nil
}

// Rewrites query string. Replaces field names
// like "Age" to "Payload.Age"
func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1:")
}
