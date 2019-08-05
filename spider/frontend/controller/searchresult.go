package controller

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"spider/engine"
	"spider/frontend/model"
	"spider/frontend/view"
	"strconv"
	"strings"

	"gopkg.in/olivere/elastic.v6"
)

// SearchResultHandler ;
type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

// CreateSearchResultHandler ;
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

func (h SearchResultHandler) ServeHTTP(
	w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))

	if err != nil {
		from = 0
	}

	page, err := h.getSearchResult(q, from)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = h.view.Render(w, page)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	result.Query = q

	resp, err := h.client.
		Search("ptt_lol").
		Query(elastic.NewQueryStringQuery(q)).
		From(from).
		Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return result, err
	}

	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	result.PreFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)

	//fmt.Println(result.Items)

	return result, nil
}
