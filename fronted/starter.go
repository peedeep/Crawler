package main

import (
	"crawler/fronted/controller"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("fronted/view")))
	repoFrontend := "fronted/view/template.html"
	http.Handle("/search", controller.CreateSearchResultHandler(repoFrontend))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}