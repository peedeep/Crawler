package zhenai

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	//contents, err := fetcher.Fetch("https://www.zhenai.com/zhenghun")
	contents, err := ioutil.ReadFile("cityList_test_data.html")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s", contents)
	result := ParseCityList(contents, "")
	const resultSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string{
		"阿坝1",
		"阿克苏1",
		"阿拉善盟1",
	}
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d reqeusts; but had %d", resultSize, len(result.Requests))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s", i, url, result.Requests[i].Url)
		}
	}
	for i, city := range expectedCities {
		if result.Items[i].Id != city {
			t.Errorf("expected city #%d: %s; but was %s", i, city, result.Items[i].Id)
		}
	}
}
