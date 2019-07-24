package fetcher

import (
	"fmt"
	"net/http"
	"testing"
)

func TestFetch(t *testing.T) {
	//http://www.zhenai.com/zhenghun/anshan/zhonglaonian
	//http://album.zhenai.com/u/1553421505
	contents, err := Fetch("http://album.zhenai.com/u/1553421505")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", contents)
}

func TestFetch2(t *testing.T) {
	resp, err := http.Get("http://album.zhenai.com/u/1553421505")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", resp.Body)
}
