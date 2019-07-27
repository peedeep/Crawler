package main

import (
	"crawler/fronted/controller"
	"net/http"
	"testing"
)

func TestStarter(t *testing.T) {
	http.Handle("/search", controller.CreateSearchResultHandler(
		`fronted\view\template.html`))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}